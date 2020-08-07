package models

import (
	"testing"
)

func testMap() *Map {
	result := NewMap("Test", 1000, 1000)
	gold := NewResourceType("Gold", 100, 10)
	result.AddResourceType(gold)
	result.AddFactoryType(NewFactoryType("GoldMine", gold.GetID(), 0.2))
	port1 := NewPortType("Port1", NewPosition(200, 100))
	result.AddPortType(port1)
	return result
}

func TestShouldIncreaseResourcesInFactory(t *testing.T) {
	currentmap := testMap()
	factory := NewFactory(*currentmap, "GoldMine", "Port1", 0.7, 0, "TestPlayer")
	factory.UpdateStorage(*currentmap, 30)
	t.Log(factory.Storage.String() + "\n")
	if factory.Storage.Amount <= 0 {
		t.Error("Expected amount to have increased")
	}
}
