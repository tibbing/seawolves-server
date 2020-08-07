package main

import (
	"apigw"
	"dynamodb"
	"errors"
	"fmt"
	"maps"
	"models"
	"strings"
	"testing"
)

func getNewMockedGame() models.Game {
	currentmap, _ := maps.GetMapByID("Scandinavia")
	player := models.NewPlayer("Player1", "Player1", models.Human)
	game := *models.NewGame(currentmap.GetID(), []models.Player{*player})
	port := models.NewPort("Stockholm", "")
	factory := models.NewFactory(currentmap, "GoldMine", "Stockholm", 0.7, 0, player.GetID())
	port.AddFactory(*factory)
	game.AddPort(port)
	return game
}

func makeRequest(event UpdateGameEvent, mockGameState models.Game) (UpdateGameResponse, error) {
	var responseTyped UpdateGameResponse
	req := apigw.GetTestRequest(event, "Player1")
	dependencies := &Dependencies{
		dynamodbClient: dynamodb.DBInstance{
			Client: dynamodb.MockedClient{MockedResponse: dynamodb.ToGetItemOutput(mockGameState)},
		},
	}

	response, err := CreateHandler(dependencies)(req)
	if err != nil {
		return responseTyped, err
	}

	responseTyped, ok := response.(UpdateGameResponse)
	if ok == false {
		return responseTyped, errors.New("Invalid response")
	}

	return responseTyped, nil
}
func TestUpdateDay1(t *testing.T) {
	event := UpdateGameEvent{
		GameID: "TestGame",
		Day:    1,
	}
	response, err := makeRequest(event, getNewMockedGame())
	if err != nil {
		t.Errorf("Request failed (day 1): %s", err.Error())
		return
	}

	if response.Game.Ports["Stockholm"].Factories[0].Storage.Amount > 1 {
		t.Errorf("Expected storage to be almost 0, but was %f", response.Game.Ports["Stockholm"].Factories[0].Storage.Amount)
		return
	}
}

func TestUpdateDay30(t *testing.T) {
	event := UpdateGameEvent{
		GameID: "TestGame",
		Day:    30,
	}
	response, err := makeRequest(event, getNewMockedGame())
	if err != nil {
		t.Errorf("Request failed (day 30): %s", err.Error())
		return
	}

	if response.Game.Ports["Stockholm"].Factories[0].Storage.Amount <= 0 {
		t.Error("Expected storage to be greater than 0")
		return
	}
}

func TestUpdatePreviousDay(t *testing.T) {
	event := UpdateGameEvent{
		GameID: "TestGame",
		Day:    20,
	}

	gameState := getNewMockedGame()
	gameState.Players["Player1"].Day = 30

	_, err := makeRequest(event, gameState)
	if err == nil || !strings.Contains(fmt.Sprintf("%s", err), "Invalid day") {
		t.Error("Expected request to fail")
		return
	}
}
