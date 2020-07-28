package app

import (
	"maps"
	"models"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )
func main() {
	currentmap := maps.Scandinavia()
	player := models.NewPlayer("Player1", models.Human)
	game := models.NewGame(currentmap.GetID(), []models.Player{*player})
	factory := models.NewFactory(*currentmap, "GoldMine", "Stockholm", 0.7, player.GetID())
	print(factory.PortID + "\n")
	for k := range currentmap.GetFactoryTypes() {
		print(currentmap.GetFactoryTypes()[k] + "\n")
	}

	factory.UpdateStorage(*currentmap, 30)
	print(game.String() + "\n")
	print(factory.Storage.String() + "\n")
	print("\n\n")

}
