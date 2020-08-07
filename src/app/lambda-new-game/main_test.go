package main

import (
	"apigw"
	"dynamodb"
	"errors"
	"testing"
)

func makeRequest(event NewGameEvent) (NewGameResponse, error) {
	var responseTyped NewGameResponse
	req := apigw.GetTestRequest(event, "Player1")
	dependencies := &Dependencies{
		dynamodbClient: dynamodb.DBInstance{
			Client: dynamodb.MockedClient{},
		},
	}

	response, err := CreateHandler(dependencies)(req)
	if err != nil {
		return responseTyped, err
	}

	responseTyped, ok := response.(NewGameResponse)
	if ok == false {
		return responseTyped, errors.New("Invalid response")
	}

	return responseTyped, nil
}

func TestCreateNewGame(t *testing.T) {
	event := NewGameEvent{
		MapID:      "Scandinavia",
		PlayerName: "Player1",
		NumAI:      2,
	}
	response, err := makeRequest(event)
	if err != nil {
		t.Errorf("Request failed: %s", err.Error())
		return
	}

	if len(response.Game.Players) != 3 {
		t.Errorf("Expected 3 players, but got %v", len(response.Game.Players))
		return
	}
}
