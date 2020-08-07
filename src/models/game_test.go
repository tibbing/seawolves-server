package models

import (
	"testing"
)

func TestShouldStartNewGame(t *testing.T) {
	player1 := NewPlayer("Player1", "Player1", Human)
	player2 := NewPlayer("Player2", "Player2", Human)
	game := NewGame("TestMap", []Player{*player1, *player2})
	if game.Turn != player1.GetID() {
		t.Error("Expected it to be player 1's turn")
	}
	if len(game.Players) != 2 {
		t.Error("Expected it to be 2 players, but found " + string(len(game.Players)))
	}
}
