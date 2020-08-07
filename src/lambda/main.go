package lambda

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

// GetTestRequest returns a valid API Gateway request with given event as body
func GetTestRequest(event interface{}) events.APIGatewayProxyRequest {
	jsonFile, err := os.Open("api-request.json")
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
