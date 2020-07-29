package http

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("http-handler")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x} ▶ %{message}`,
)

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

// Decorate handles errors and returns a well formed API Gateway Response
func Decorate(handler func(request events.APIGatewayProxyRequest) (interface{}, error)) func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var response events.APIGatewayProxyResponse
		log.Debugf("Processing request ID %s, Body: (%d) %s", request.RequestContext.RequestID, len(request.Body), request.Body)

		headers := map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Authorization,Content-Type",
			"Access-Control-Allow-Methods": "OPTIONS,HEAD,GET,POST,PUT,PATCH,DELETE",
			"Access-Control-Max-Age":       "7200",
		}

		if request.HTTPMethod == "get" {
			headers["Cache-Control"] = "private, max-age=300"
		} else {
			headers["Cache-Control"] = "no-store"
		}

		result, err := handler(request)
		if err != nil {
			response = errorResponse(err, 500, headers)
		}

		responseJSON, err := json.Marshal(result)
		if err != nil {
			response = errorResponse(err, 500, headers)
		}

		response = events.APIGatewayProxyResponse{
			StatusCode:      200,
			Body:            string(responseJSON),
			Headers:         headers,
			IsBase64Encoded: false,
		}
		return response, nil
	}

}

func errorResponse(err error, code int, headers map[string]string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:      code,
		Body:            err.Error(),
		Headers:         headers,
		IsBase64Encoded: false,
	}
}
