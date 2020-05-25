package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/smartatransit/third_rail/pkg/models"
	log "github.com/sirupsen/logrus"
)

func Seed(db *gorm.DB) {

	db.LogMode(true)

	//Directions
	log.Info("Creating Directions")

	var northbound models.Direction
	db.FirstOrCreate(&northbound, &models.Direction{
		Name: "Northbound",
	})

	var southbound models.Direction
	db.FirstOrCreate(&southbound, &models.Direction{
		Name: "Southbound",
	})

	var eastbound models.Direction
	db.FirstOrCreate(&eastbound, &models.Direction{
		Name: "Eastbound",
	})

	var westbound models.Direction
	db.FirstOrCreate(&westbound, &models.Direction{
		Name: "Westbound",
	})


	//Lines
	log.Info("Creating Lines")

	var gold models.Line
	db.FirstOrCreate(&gold, &models.Line{
		Name:       "Gold",
	})
	db.Model(&gold).Association("Directions").Append([]models.Direction{northbound, southbound})

	var red models.Line
	db.FirstOrCreate(&red, &models.Line{
		Name:       "Red",
	})
	db.Model(&gold).Association("Directions").Append([]models.Direction{northbound, southbound})

	var blue models.Line
	db.FirstOrCreate(&blue, &models.Line{
		Name:       "Blue",
	})
	db.Model(&gold).Association("Directions").Append([]models.Direction{eastbound, westbound})

	var green models.Line
	db.FirstOrCreate(&green, &models.Line{
		Name:       "Green",
	})
	db.Model(&gold).Association("Directions").Append([]models.Direction{eastbound, westbound})

	//Stations
	log.Info("Creating Stations")

	var doraville models.Station
	db.FirstOrCreate(&doraville, &models.Station{
		Name:  "Doraville",
	})
	db.Model(&doraville).Association("Lines").Append([]models.Line{gold})

	var doravilleDetail models.StationDetail
	db.FirstOrCreate(&doravilleDetail, &models.StationDetail{
		StationID: doraville.ID,
		Description: "Doraville station",
		Location: "dnh0f5v6mxzj",
	})
}