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
	"math"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("purchase:factory")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x} ▶ %{message}`,
)

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

// PurchaseFactoryEvent lambda event
type PurchaseFactoryEvent struct {
	GameID        string
	PortID        string
	FactoryTypeID string
	LocationID    int
}

// PurchaseFactoryResponse lambda response
type PurchaseFactoryResponse struct {
	Game *models.Game
}

// Dependencies Lambda dependencies
type Dependencies struct {
	dynamodbClient dynamodb.DBInstance
}

// CreateHandler creates the event handler
func CreateHandler(dependencies *Dependencies) func(request events.APIGatewayProxyRequest) (interface{}, error) {
	return func(request events.APIGatewayProxyRequest) (interface{}, error) {
		var event PurchaseFactoryEvent
		json.Unmarshal([]byte(request.Body), &event)

		response := PurchaseFactoryResponse{Game: nil}

		game, err := dependencies.dynamodbClient.GetGameByID(event.GameID)
		if err != nil {
			return response, err
		}

		playerID, err := apigw.GetUserID(request)
		if err != nil {
			return PurchaseFactoryResponse{Game: nil}, err
		}

		if game.ID == "" {
			return PurchaseFactoryResponse{Game: nil}, fmt.Errorf("Invalid game: %s", event.GameID)
		}
		currentMap, err := maps.GetMapByID(game.MapID)
		if err != nil {
			return response, err
		}

		port := game.Ports[event.PortID]
		if port == nil {
			return PurchaseFactoryResponse{Game: nil}, fmt.Errorf("Invalid port: %s", event.PortID)
		}

		portType := currentMap.PortTypes[port.PortTypeID]

		priceModifier := (float64(len(port.Factories)) / float64(portType.FactoryLocations)) + portType.FactoryPriceModifier
		price := float64(currentMap.FactoryTypes[event.FactoryTypeID].BasePrice) * priceModifier
		game.Players[playerID].MakeTransaction(-int(math.Round(price)))

		dependencies.dynamodbClient.UpdateGame(game)

		return PurchaseFactoryResponse{Game: &game}, nil
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
