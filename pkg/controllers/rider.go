package controllers

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/smartatransit/third_rail/pkg/models"
	"net/http"
)

type RiderController struct{}

// GetAlerts godoc
// @Summary Get Alerts from various MARTA sources
// @Description MARTA alerts sourced from their official twitter account
// @Produce  json
// @Success 200 {object} alertResponse
// @Router /live/alerts [get]
// @Security ApiKeyAuth
func (controller RiderController) GetRiderAlerts(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var station models.Station
	db.Preload("Info").First(&station, "Name = ?", "Midtown Station") // find product with code l1212

	json.NewEncoder(w).Encode(station)
}

func (controller RiderController) Migrate(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	// Create
	/*db.Create(&models.Station{
		Name: "Midtown Station",
		Info: models.StationInfo{
			Description: "Midtown Station",
			Location:    "",
		}})

	log.Info("Created data")

	// Read
	var station models.Station
	db.First(&station, 1)                             // find product with id 1
	db.First(&station, "Name = ?", "Midtown Station") // find product with code l1212

	log.Infof("Read data: %v", &station)

	// Update - update product's price to 2000
	db.Model(&station.Info).Update("Description", "Midtown Station in Atlanta")

	log.Info("Updated data")

	// Delete - delete product
	//db.Delete(&station)

	//log.Info("Deleted data")*/
}
