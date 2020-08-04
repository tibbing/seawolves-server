package dynamodb

import (
	"encoding/json"
	"maps"
	"models"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type MockedClient struct {
	dynamodbiface.DynamoDBAPI
	Resp dynamodb.GetItemOutput
}

func getNewMockedGame() models.Game {
	currentmap := maps.Scandinavia()
	player := models.NewPlayer("Player1", models.Human)
	game := *models.NewGame(currentmap.GetID(), []models.Player{*player})
	port := models.NewPort("Stockholm", "")
	factory := models.NewFactory(*currentmap, "GoldMine", "Stockholm", 0.7, 0, player.GetID())
	port.AddFactory(*factory)
	game.AddPort(port)
	return game
}

// GetMockedGameOutput returns mocked game as dynamodb GetItem output
func GetMockedGameOutput() dynamodb.GetItemOutput {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(getNewMockedGame())
	log.Debug(string(inrec) + "\n")
	json.Unmarshal(inrec, &inInterface)

	output := dynamodb.GetItemOutput{}
	item, _ := dynamodbattribute.MarshalMap(inInterface)
	output.SetItem(item)
	return output
}

// GetItem mocked implementation that returns pre-defined response
func (m MockedClient) GetItem(in *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	// if _, ok := in.Key["GameID"]; ok {
	// 	// Getting game by ID
	// 	gameID := in.Key["GameID"].String()
	// 	var game models.Game
	// 	if _, gameIDExists := m.games[gameID]; gameIDExists {
	// 		log.Infof("Game '%s' exists", gameID)
	// 		game = m.games[gameID]
	// 	} else {
	// 		log.Infof("Game '%s' does not exist, creating new...", gameID)
	// 		if m.games == nil {
	// 			m.games = map[string]models.Game{}
	// 		}
	// 		game = getNewMockedGame()
	// 		m.games[gameID] = game
	// 	}

	// 	// Set Game as DynamoDB output
	// 	item := &m.Resp
	// 	item.SetItem(map[string]*dynamodb.AttributeValue{})
	// 	return item, nil
	// }
	return &m.Resp, nil
}
