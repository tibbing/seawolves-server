package main

import (
	"context"
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

// HandleLambdaEvent lambda handler
func createHandler(request events.APIGatewayProxyRequest) (interface{}, error) {
	var event UpdateGameEvent
	json.Unmarshal([]byte(request.Body), &event)

	game, err := getGameByID(event.GameID)
	if err != nil {
		return UpdateGameResponse{Game: nil}, err
	}
	currentmap := maps.Scandinavia()
	game.UpdateForPlayer(currentmap, event.PlayerID, event.Day)

	return UpdateGameResponse{Game: &game}, nil
}

func getGameByID(gameID string) (models.Game, error) {
	currentmap := maps.Scandinavia()
	player := models.NewPlayer("Player1", models.Human)
	game := models.NewGame(currentmap.GetID(), []models.Player{*player})
	port := models.NewPort("Stockholm", "")
	factory := models.NewFactory(*currentmap, "GoldMine", "Stockholm", 0.7, 0, player.GetID())
	port.AddFactory(*factory)
	game.AddPort(port)
	return *game, nil
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return http.Decorate(createHandler)(ctx, request)
}
