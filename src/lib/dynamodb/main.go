package dynamodb

import (
	"encoding/json"
	"errors"
	"fmt"
	"lib/models"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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

// GetClient returns a DynamoDB client
func GetClient(config *aws.Config) *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(config))
	client := dynamodb.New(sess)
	return client
}

// UpdateGame updates game state in DynamoDB
func (q *DBInstance) UpdateGame(game models.Game) error {

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(game)
	log.Debug(string(inrec) + "\n")
	json.Unmarshal(inrec, &inInterface)
	updateExpr := buildUpdateExpression(inInterface, "", "")
	exprAttrValues := buildExpressionAttributeValues(inInterface, "")
	// exprAttrValues, _ := dynamodbattribute.MarshalMap(inInterface)

	// fmt.Println(fmt.Sprintf("ExpressionAttributeValues %v", exprAttrValues))
	// fmt.Println(fmt.Sprintf("updateExpr %s", updateExpr))

	params := dynamodb.UpdateItemInput{
		TableName:                 aws.String("Games"),
		UpdateExpression:          &updateExpr,
		ExpressionAttributeValues: exprAttrValues,
		ReturnValues:              aws.String("UPDATED_NEW"),
		Key: map[string]*dynamodb.AttributeValue{
			"GameID": {
				S: aws.String(game.GetID()),
			},
		},
	}

	_, err := q.Client.UpdateItem(&params)
	return err
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
