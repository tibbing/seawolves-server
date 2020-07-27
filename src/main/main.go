package main

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
	factory := models.NewFactory(*currentmap, "GoldMine", "Stockholm", 0.7, player.GetID())
	print(factory.PortID + "\n")
	for k := range currentmap.GetFactoryTypes() {
		print(currentmap.GetFactoryTypes()[k] + "\n")
	}

	factory.UpdateStorage(*currentmap, 30)
	print(factory.Storage.Amount)
	print("\n\n")

}
