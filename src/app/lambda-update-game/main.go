package lambda-update-game

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	PlayerID string
	GameID  string
	Day  int16
}

type Response struct {
	Message string
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
