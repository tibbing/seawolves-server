package main

import (
	"dynamodb"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"models"
	"os"
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func getTestRequest(event UpdateGameEvent) events.APIGatewayProxyRequest {
	jsonFile, err := os.Open("request.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var req events.APIGatewayProxyRequest
	json.Unmarshal(byteValue, &req)
	eventBody, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
	}
	req.Body = string(eventBody)

	return req
}

func makeRequest(event UpdateGameEvent, mockGameState models.Game) (UpdateGameResponse, error) {
	var responseTyped UpdateGameResponse
	req := getTestRequest(event)
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
		GameID:   "TestGame",
		PlayerID: "Player1",
		Day:      1,
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
		GameID:   "TestGame",
		PlayerID: "Player1",
		Day:      30,
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
		GameID:   "TestGame",
		PlayerID: "Player1",
		Day:      20,
	}

	gameState := getNewMockedGame()
	gameState.Players["Player1"].Day = 30

	_, err := makeRequest(event, gameState)
	if err == nil || !strings.Contains(fmt.Sprintf("%s", err), "Invalid day") {
		t.Error("Expected request to fail")
		return
	}
}
