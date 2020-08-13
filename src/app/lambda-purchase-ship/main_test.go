package main

import (
	"errors"
	"fmt"
	"lib/apigw"
	"lib/dynamodb"
	"lib/maps"
	"lib/models"
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

	shipyard := models.NewShipyard(currentmap, 3, []string{"Brig", "Corvette", "Fluyt", "Galleon"})
	port.CreateShipyard(*shipyard)

	game.AddPort(port)
	return game
}

func makeRequest(event PurchaseShipEvent, mockGameState models.Game) (PurchaseShipResponse, error) {
	var responseTyped PurchaseShipResponse
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

	responseTyped, ok := response.(PurchaseShipResponse)
	if ok == false {
		return responseTyped, errors.New("Invalid response")
	}

	return responseTyped, nil
}

func TestPurchaseShip(t *testing.T) {
	event := PurchaseShipEvent{
		GameID:     "TestGame",
		PortID:     "Stockholm",
		ShipTypeID: "Brig",
	}

	response, err := makeRequest(event, getNewMockedGame())
	if err != nil {
		t.Errorf("Request failed: %s", err.Error())
		return
	}

	fleets := response.Game.Players["Player1"].Fleets
	if len(fleets) != 1 {
		t.Errorf("No fleets found")
		return
	}

	if len(fleets[0].Ships) != 1 {
		t.Errorf("No ships found")
		return
	}

	goldBefore := getNewMockedGame().Players["Player1"].Gold
	goldAfter := response.Game.Players["Player1"].Gold

	if goldAfter >= goldBefore {
		t.Errorf("Gold has not been deducted")
		return
	}

	if response.Game.Ports["Stockholm"].Shipyard.HasShip("Brig") {
		t.Errorf("Expected ship to be removed from shipyard")
		return
	}
}

func TestPurchaseInvalidPort(t *testing.T) {
	event := PurchaseShipEvent{
		GameID:     "TestGame",
		PortID:     "Invalid",
		ShipTypeID: "Brig",
	}

	_, err := makeRequest(event, getNewMockedGame())
	if err == nil || !strings.Contains(fmt.Sprintf("%s", err), "Invalid port") {
		t.Error("Expected request to fail")
		return
	}
}
