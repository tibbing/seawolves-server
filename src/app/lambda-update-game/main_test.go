package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func getTestEvent() events.APIGatewayProxyRequest {
	jsonFile, err := os.Open("valid_event.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var req events.APIGatewayProxyRequest
	json.Unmarshal(byteValue, &req)
	return req
}

func TestValidEvent(t *testing.T) {
	testEvent := getTestEvent()
	response, err := CreateHandler(testEvent)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(response)

	responseTyped, ok := response.(UpdateGameResponse)
	if ok == false {
		t.Error("Invalid response")
		return
	}
	t.Log(responseTyped.Game.Ports["Stockholm"].Factories[0].Storage.String())
}
