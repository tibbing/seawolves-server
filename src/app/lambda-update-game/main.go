package lambda

import (
	"maps"
	"models"

	"github.com/aws/aws-lambda-go/lambda"
)

// UpdateGameEvent lambda event
type UpdateGameEvent struct {
	PlayerID string
	GameID   string
	Day      int16
}

// UpdateGameResponse lambda response
type UpdateGameResponse struct {
	Game string
}

// HandleLambdaEvent lambda handler
func HandleLambdaEvent(event UpdateGameEvent) (UpdateGameResponse, error) {
	game, err := getGameByID(event.GameID)
	if err != nil {
		return UpdateGameResponse{Game: ""}, err
	}
	currentmap := maps.Scandinavia()
	game.UpdateForPlayer(currentmap, event.PlayerID, event.Day)
	return UpdateGameResponse{Game: "test"}, nil
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
	lambda.Start(HandleLambdaEvent)
}
