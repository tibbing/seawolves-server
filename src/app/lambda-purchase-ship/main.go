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

var log = logging.MustGetLogger("purchase:ship")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x} ▶ %{message}`,
)

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

// PurchaseShipEvent lambda event
type PurchaseShipEvent struct {
	GameID     string
	ShipTypeID string
	PortID     string
}

// PurchaseShipResponse lambda response
type PurchaseShipResponse struct {
	Game *models.Game
}

// Dependencies Lambda dependencies
type Dependencies struct {
	dynamodbClient dynamodb.DBInstance
}

// CreateHandler creates the event handler
func CreateHandler(dependencies *Dependencies) func(request events.APIGatewayProxyRequest) (interface{}, error) {
	return func(request events.APIGatewayProxyRequest) (interface{}, error) {
		var event PurchaseShipEvent
		json.Unmarshal([]byte(request.Body), &event)

		response := PurchaseShipResponse{Game: nil}

		game, err := dependencies.dynamodbClient.GetGameByID(event.GameID)
		if err != nil {
			return response, err
		}

		playerID, err := apigw.GetUserID(request)
		if err != nil {
			return PurchaseShipResponse{Game: nil}, err
		}

		if game.ID == "" {
			return PurchaseShipResponse{Game: nil}, fmt.Errorf("Invalid game: %s", event.GameID)
		}
		currentMap, err := maps.GetMapByID(game.MapID)
		if err != nil {
			return response, err
		}

		port := game.Ports[event.PortID]
		if port == nil {
			return PurchaseShipResponse{Game: nil}, fmt.Errorf("Invalid port: %s", event.PortID)
		}

		shipyard := port.Shipyard
		if shipyard == nil {
			return PurchaseShipResponse{Game: nil}, fmt.Errorf("No shipyard in port %s", event.PortID)
		}

		if !shipyard.HasShip(event.ShipTypeID) {
			return PurchaseShipResponse{Game: nil}, fmt.Errorf("Ship type %s does not exist in port %s", event.ShipTypeID, event.PortID)
		}

		shipyard.RemoveShip(event.ShipTypeID)
		ship := models.NewShip("New ship", event.ShipTypeID)
		fleet := models.NewFleet("New fleet", playerID, []models.Ship{*ship})
		game.Players[playerID].MakeTransaction(-currentMap.ShipTypes[event.ShipTypeID].Price)
		game.Players[playerID].AddFleet(*fleet)

		dependencies.dynamodbClient.UpdateGame(game)

		return PurchaseShipResponse{Game: &game}, nil
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
