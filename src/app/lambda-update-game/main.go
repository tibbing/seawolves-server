package main

import (
	"context"
	"dynamodb"
	"encoding/json"
	"http-handler"
	"maps"
	"models"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
	PlayerID string
	GameID   string
	Day      int16
}

// UpdateGameResponse lambda response
type UpdateGameResponse struct {
	Game *models.Game
}

// CreateHandler creates the event handler
func CreateHandler(request events.APIGatewayProxyRequest) (interface{}, error) {
	var event UpdateGameEvent
	json.Unmarshal([]byte(request.Body), &event)

	game, err := dynamodb.GetGameByID(event.GameID)
	if err != nil {
		return UpdateGameResponse{Game: nil}, err
	}
	currentmap := maps.Scandinavia()
	game.UpdateForPlayer(currentmap, event.PlayerID, event.Day)
	log.Info(game.Ports["Stockholm"].Factories[0].Storage.String())

	return UpdateGameResponse{Game: &game}, nil
}

func main() {
	lambda.Start(handler)
}

// Handler method
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return http.Decorate(CreateHandler)(ctx, request)
}
