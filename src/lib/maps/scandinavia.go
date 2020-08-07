package maps

import "lib/models"

// Scandinavia Creates an instance of the map Scandinavia
func Scandinavia() *models.Map {
	result := models.NewMap("Scandinavia", 1000, 1000)

	gold := models.NewResourceType("Gold", 100, 10)
	result.AddResourceType(gold)
	silver := models.NewResourceType("Silver", 50, 10)
	result.AddResourceType(silver)

	result.AddFactoryType(models.NewFactoryType("GoldMine", gold.GetID(), 0.2))
	result.AddFactoryType(models.NewFactoryType("SilverMine", silver.GetID(), 0.3))

	stockholm := models.NewPortType("Stockholm", models.NewPosition(200, 100))
	result.AddPortType(stockholm)
	goteborg := models.NewPortType("Göteborg", models.NewPosition(140, 100))
	result.AddPortType(goteborg)

	return result
}
