package dynamodb

import (
	"maps"
	"models"
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("http-handler")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x} ▶ %{message}`,
)

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

// GetGameByID returns game by ID
func GetGameByID(gameID string) (models.Game, error) {
	currentmap := maps.Scandinavia()
	player := models.NewPlayer("Player1", models.Human)
	game := models.NewGame(currentmap.GetID(), []models.Player{*player})
	port := models.NewPort("Stockholm", "")
	factory := models.NewFactory(*currentmap, "GoldMine", "Stockholm", 0.7, 0, player.GetID())
	port.AddFactory(*factory)
	game.AddPort(port)
	return *game, nil
}
