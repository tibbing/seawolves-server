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
	factory := models.NewFactory("GoldMine", "Stockholm", player.GetID())
	print(factory.PortID + "\n")
	for k := range currentmap.GetFactoryTypes() {
		print(currentmap.GetFactoryTypes()[k] + "\n")
	}
}
