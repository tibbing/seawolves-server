package apigw

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aws/aws-lambda-go/events"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("apigw:handler")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x} ▶ %{message}`,
)

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

// GetTestRequest returns a valid API Gateway request with given event as body
func GetTestRequest(event interface{}, userID string) events.APIGatewayProxyRequest {
	absPath, _ := filepath.Abs("/go/src/lib/apigw/request.json")
	jsonFile, err := os.Open(absPath)
	if err != nil {
		log.Error(err)
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

	claims := make(map[string]interface{})
	claims["sub"] = userID
	authorizer := make(map[string]interface{})
	authorizer["claims"] = claims
	req.RequestContext.Authorizer = authorizer

	return req
}

// GetUserID Returns user ID from request claims
func GetUserID(req events.APIGatewayProxyRequest) (string, error) {
	if claims, ok := req.RequestContext.Authorizer["claims"]; ok {
		v, ok := claims.(map[string]interface{})
		if !ok {
			log.Error(claims)
			return "", errors.New("Invalid claims")
		}
		return v["sub"].(string), nil
	}
	return "", errors.New("Claims have not been set")
}
