package dynamodb

import (
	"maps"
	"models"
)

// MockClient mocked dynamodb client
type MockClient struct {
	games map[string]models.Game
}

// GetGameByID gets mocked game by ID
func (m *MockClient) GetGameByID(gameID string) (models.Game, error) {

	if game, ok := m.games[gameID]; ok {
		log.Infof("Game '%s' exists", gameID)
		return game, nil
	}
	log.Infof("Game '%s' does not exist, creating new...", gameID)
	if m.games == nil {
		m.games = map[string]models.Game{}
	}
	currentmap := maps.Scandinavia()
	player := models.NewPlayer("Player1", models.Human)
	game := models.NewGame(currentmap.GetID(), []models.Player{*player})
	port := models.NewPort("Stockholm", "")
	factory := models.NewFactory(*currentmap, "GoldMine", "Stockholm", 0.7, 0, player.GetID())
	port.AddFactory(*factory)
	game.AddPort(port)
	m.games[gameID] = *game
	return *game, nil
}
