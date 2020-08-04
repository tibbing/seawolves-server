package dynamodb

import (
	"errors"
	"fmt"
	"models"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
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

// DBInstance the DynamoDB instance
type DBInstance struct {
	Client dynamodbiface.DynamoDBAPI
}

// GetGameByID returns a game by ID
func (q *DBInstance) GetGameByID(gameID string) (models.Game, error) {
	var game models.Game

	params := dynamodb.GetItemInput{
		TableName: aws.String("Games"),
		Key: map[string]*dynamodb.AttributeValue{
			"GameID": {
				S: aws.String(gameID),
			},
		},
	}

	result, err := q.Client.GetItem(&params)
	if err != nil {
		log.Error(err.Error())
		return game, err
	}
	if result.Item == nil {
		return game, errors.New("Could not find gameID '" + gameID + "'")
	}

	log.Debugf("Found game in DB")

	game = models.Game{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &game)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	// fmt.Printf("%s\n", result.Item)
	// fmt.Printf("%s\n", game.String())

	return game, nil
}
