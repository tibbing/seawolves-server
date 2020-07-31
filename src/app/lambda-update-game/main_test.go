package main

import (
	"dynamodb"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

var dependencies *Dependencies

func makeRequest(event UpdateGameEvent) (UpdateGameResponse, error) {
	var responseTyped UpdateGameResponse
	req := getTestRequest(event)
	if dependencies == nil {
		dependencies = &Dependencies{
			dynamodbClient: &dynamodb.MockClient{},
		}
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
func TestValidEvent(t *testing.T) {
	event := UpdateGameEvent{
		GameID:   "TestGame",
		PlayerID: "Player1",
		Day:      0,
	}
	response, err := makeRequest(event)
	if err != nil {
		t.Errorf("Request failed (day 0): %s", err.Error())
	}

	if response.Game.Ports["Stockholm"].Factories[0].Storage.Amount != 0 {
		t.Errorf("Expected storage to be 0, but was %f", response.Game.Ports["Stockholm"].Factories[0].Storage.Amount)
	}

	event = UpdateGameEvent{
		GameID:   "TestGame",
		PlayerID: "Player1",
		Day:      30,
	}
	response, err = makeRequest(event)
	if err != nil {
		t.Errorf("Request failed (day 30): %s", err.Error())
	}

	if response.Game.Ports["Stockholm"].Factories[0].Storage.Amount <= 0 {
		t.Error("Expected storage to be greater than 0")
	}
}
