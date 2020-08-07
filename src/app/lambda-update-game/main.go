package main

import (
	"context"
	"encoding/json"
	"errors"
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

var log = logging.MustGetLogger("update:game")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x} ▶ %{message}`,
)

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

// UpdateGameEvent lambda event
type UpdateGameEvent struct {
	GameID string
	Day    int
}

// UpdateGameResponse lambda response
type UpdateGameResponse struct {
	Game *models.Game
}

// Dependencies Lambda dependencies
type Dependencies struct {
	dynamodbClient dynamodb.DBInstance
}

// CreateHandler creates the event handler
func CreateHandler(dependencies *Dependencies) func(request events.APIGatewayProxyRequest) (interface{}, error) {
	return func(request events.APIGatewayProxyRequest) (interface{}, error) {
		var event UpdateGameEvent
		json.Unmarshal([]byte(request.Body), &event)

		game, err := dependencies.dynamodbClient.GetGameByID(event.GameID)
		if err != nil {
			return UpdateGameResponse{Game: nil}, err
		}
		if game.ID == "" {
			return UpdateGameResponse{Game: nil}, errors.New("Invalid game: " + event.GameID)
		}
		currentmap, err := maps.GetMapByID(game.MapID)
		if err != nil {
			return UpdateGameResponse{Game: nil}, err
		}

		playerID, err := apigw.GetUserID(request)
		if err != nil {
			return UpdateGameResponse{Game: nil}, err
		}

		err = game.UpdateForPlayer(&currentmap, playerID, event.Day)
		if err != nil {
			return UpdateGameResponse{Game: nil}, err
		}

		dependencies.dynamodbClient.UpdateGame(game)

		return UpdateGameResponse{Game: &game}, nil
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
