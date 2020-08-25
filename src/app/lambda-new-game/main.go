package main

import (
	"context"
	"encoding/json"
	"fmt"
	"lib/apigw"
	"lib/dynamodb"
	"lib/http-handler"
	"lib/maps"
	"lib/models"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("new:game")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x} ▶ %{message}`,
)

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

// NewGameEvent lambda event
type NewGameEvent struct {
	PlayerName string
	MapID      string
	NumAI      int
}

// NewGameResponse lambda response
type NewGameResponse struct {
	Game *models.Game
}

// Dependencies Lambda dependencies
type Dependencies struct {
	dynamodbClient dynamodb.DBInstance
}

// CreateHandler creates the event handler
func CreateHandler(dependencies *Dependencies) func(request events.APIGatewayProxyRequest) (interface{}, error) {
	return func(request events.APIGatewayProxyRequest) (interface{}, error) {
		var event NewGameEvent
		json.Unmarshal([]byte(request.Body), &event)

		response := NewGameResponse{Game: nil}

		currentmap, err := maps.GetMapByID(event.MapID)
		if err != nil {
			return response, err
		}

		if event.NumAI < 0 || event.NumAI > 3 {
			return response, fmt.Errorf("Invalid number of AI players: %v", event.NumAI)
		}

		if event.PlayerName == "" || len(event.PlayerName) > 20 {
			return response, fmt.Errorf("Invalid player name: %v", event.PlayerName)
		}

		userID, err := apigw.GetUserID(request)
		if err != nil {
			return response, err
		}

		players := make([]models.Player, 1+event.NumAI)
		players[0] = *models.NewPlayer(event.PlayerName, userID, models.Human)
		for i := 0; i < event.NumAI; i++ {
			players[i+1] = *models.NewPlayer(fmt.Sprintf("Ai%v", i), fmt.Sprintf("Ai%v", i), models.AI)
		}
		game := *models.NewGame(currentmap.GetID(), players)

		dependencies.dynamodbClient.UpdateGame(game)

		response = NewGameResponse{Game: &game}
		return response, nil
	}
}

func main() {
	lambda.Start(handler)
}

// Set up dependencies
func createDefaultHandler() func(request events.APIGatewayProxyRequest) (interface{}, error) {
	config := &aws.Config{
		Region: aws.String("eu-north-1"),
	}

	dependencies := Dependencies{
		dynamodbClient: dynamodb.DBInstance{
			Client: dynamodb.GetClient(config),
		},
	}
	return CreateHandler(&dependencies)
}

// Handler method
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return http.Decorate(createDefaultHandler())(ctx, request)
}
