package main

import (
	"context"
	"dynamodb"
	"encoding/json"
	"fmt"
	"http-handler"
	"maps"
	"models"
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
	Players []string
	MapID   string
	NumAI   int
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

		currentmap, err := maps.GetMapByID(event.MapID)
		if err != nil {
			return NewGameResponse{Game: nil}, err
		}

		if event.NumAI < 0 || event.NumAI > 3 {
			return NewGameResponse{Game: nil}, fmt.Errorf("Invalid number of AI players: %v", event.NumAI)
		}

		if len(event.Players) < 1 || len(event.Players) > 3 {
			return NewGameResponse{Game: nil}, fmt.Errorf("Invalid number of Human players: %v", len(event.Players))
		}

		players := make([]models.Player, len(event.Players)+event.NumAI)
		for i, playerID := range event.Players {
			players[i] = *models.NewPlayer(playerID, models.Human)
		}
		for i := 0; i < event.NumAI; i++ {
			players[i+len(event.Players)] = *models.NewPlayer(fmt.Sprintf("Ai%v", i), models.AI)
		}
		game := *models.NewGame(currentmap.GetID(), players)

		dependencies.dynamodbClient.UpdateGame(game)

		return NewGameResponse{Game: &game}, nil
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
