package maps

import "models"

// Scandinavia Creates an instance of the map Scandinavia
func Scandinavia() *models.Map {
	result := models.NewMap("Scandinavia", 1000, 1000)

	gold := models.NewResourceType("Gold", 100, 10)
	result.AddResourceType(gold)
	silver := models.NewResourceType("Silver", 50, 10)
	result.AddResourceType(silver)

	result.AddFactoryType(models.NewFactoryType("GoldMine", gold.GetID()))
	result.AddFactoryType(models.NewFactoryType("SilverMine", silver.GetID()))

	stockholm := models.NewPort("Stockholm", models.NewPosition(200, 100))
	result.AddPort(stockholm)
	goteborg := models.NewPort("Göteborg", models.NewPosition(140, 100))
	result.AddPort(goteborg)

	return result
}
