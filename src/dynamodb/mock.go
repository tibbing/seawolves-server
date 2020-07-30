package dynamodb

import (
	"maps"
	"models"
)

// MockClient mocked dynamodb client
type MockClient struct{}

// GetGameByID gets mocked game by ID
func (m *MockClient) GetGameByID(gameID string) (models.Game, error) {
	log.Infof("Mocking game state...")
	currentmap := maps.Scandinavia()
	player := models.NewPlayer("Player1", models.Human)
	game := models.NewGame(currentmap.GetID(), []models.Player{*player})
	port := models.NewPort("Stockholm", "")
	factory := models.NewFactory(*currentmap, "GoldMine", "Stockholm", 0.7, 0, player.GetID())
	port.AddFactory(*factory)
	game.AddPort(port)
	return *game, nil
}
