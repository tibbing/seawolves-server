package maps

import "lib/models"

// Scandinavia Creates an instance of the map Scandinavia
func Scandinavia() *models.Map {
	result := models.NewMap("Scandinavia", 1000, 1000)

	gold := models.NewResourceType("Gold", 100, 10)
	result.AddResourceType(gold)
	silver := models.NewResourceType("Silver", 50, 10)
	result.AddResourceType(silver)

	result.AddFactoryType(models.NewFactoryType("GoldMine", gold.GetID(), 0.2, 800))
	result.AddFactoryType(models.NewFactoryType("SilverMine", silver.GetID(), 0.3, 500))

	stockholm := models.NewPortType("Stockholm", models.NewPosition(200, 100), 1.2)
	result.AddPortType(stockholm)
	goteborg := models.NewPortType("Göteborg", models.NewPosition(140, 100), 1.1)
	result.AddPortType(goteborg)

	result.AddShipType(models.NewShipType("Brig", 30, 200, 80, 300))
	result.AddShipType(models.NewShipType("Fluyt", 5, 2000, 200, 600))
	result.AddShipType(models.NewShipType("Corvette", 20, 1000, 500, 1000))
	result.AddShipType(models.NewShipType("Galleon", 10, 1500, 800, 5000))

	return result
}
