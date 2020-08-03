package dynamodb

import (
	"errors"
	"fmt"
	"models"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

// DDBInstance interface
type DDBInstance interface {
	// GetGameByID() <-chan struct{}
	GetGameByID(gameID string) (models.Game, error)
}

// Client mocked dynamodb client
type Client struct {
}

// GetGameByID gets mocked game by ID
func (m *Client) GetGameByID(gameID string) (models.Game, error) {
	var game models.Game
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	tableName := "Games"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"GameID": {
				S: aws.String(gameID),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return game, err
	}
	if result.Item == nil {
		return game, errors.New("Could not find gameID '" + gameID + "'")
	}

	item := models.Game{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return game, nil
}
