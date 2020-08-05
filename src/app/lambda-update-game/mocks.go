package main

import (
	"maps"
	"models"
)

func getNewMockedGame() models.Game {
	currentmap := maps.Scandinavia()
	player := models.NewPlayer("Player1", models.Human)
	game := *models.NewGame(currentmap.GetID(), []models.Player{*player})
	port := models.NewPort("Stockholm", "")
	factory := models.NewFactory(*currentmap, "GoldMine", "Stockholm", 0.7, 0, player.GetID())
	port.AddFactory(*factory)
	game.AddPort(port)
	return game
}
