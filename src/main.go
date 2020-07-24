package main

import "models"

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )
func main() {
	test()
}

func test() {
	scandinavia := models.NewMap("Scandinavia", 1000, 1000)
	stockholm := models.NewPort("Stockholm", models.NewPosition(200, 100))
	scandinavia.AddPort(stockholm)

	gold := models.NewResourceType("Gold", 100, 10)
	goldMine := models.NewFactoryType("GoldMine", gold.GetID())

	player := models.NewPlayer("Player1", models.Human)
	factory := models.NewFactory(goldMine.GetID(), stockholm.GetID(), player.GetID())

	print(factory.PortID)
}
