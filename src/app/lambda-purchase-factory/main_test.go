package main

import (
	"errors"
	"lib/apigw"
	"lib/dynamodb"
	"lib/maps"
	"lib/models"
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

func makeRequest(event PurchaseFactoryEvent, mockGameState models.Game) (PurchaseFactoryResponse, error) {
	var responseTyped PurchaseFactoryResponse
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

	responseTyped, ok := response.(PurchaseFactoryResponse)
	if ok == false {
		return responseTyped, errors.New("Invalid response")
	}

	return responseTyped, nil
}
