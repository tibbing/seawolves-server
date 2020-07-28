package app

import (
	"encoding/json"
	"maps"
	"models"
	"testing"
)

func TestShouldStartNewGame(t *testing.T) {
	player1 := models.NewPlayer("Player1", models.Human)
	player2 := models.NewPlayer("Player1", models.Human)
	game := models.NewGame("TestMap", []models.Player{*player1, *player2})
	if game.Turn != player1.GetID() {
		t.Error("Expected it to be player 1's turn")
	}
	if len(game.Players) != 2 {
		t.Error("Expected it to be 2 players")
	}
}

func TestShouldIncreaseResourcesInFactory(t *testing.T) {
	currentmap := maps.Scandinavia()
	factory := models.NewFactory(*currentmap, "GoldMine", "Stockholm", 0.7, "TestPlayer")
	factory.UpdateStorage(*currentmap, 30)
	t.Log(factory.Storage.String() + "\n")
	if factory.Storage.Amount <= 0 {
		t.Error("Expected amount to have increased")
	}
}

func TestShouldSerializeMap(t *testing.T) {
	currentmap := maps.Scandinavia()
	e, err := json.Marshal(currentmap)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(e) + "\n")

	var result models.Map
	json.Unmarshal([]byte(e), &result)
	t.Logf("Deserialized: %s", result.String()+"\n")
}
