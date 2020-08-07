package maps

import (
	"encoding/json"
	"lib/models"
	"testing"
)

func TestShouldSerializeMap(t *testing.T) {
	currentmap := Scandinavia()
	e, err := json.Marshal(currentmap)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(e) + "\n")

	var result models.Map
	json.Unmarshal([]byte(e), &result)
	t.Logf("Deserialized: %s", result.String()+"\n")
}
