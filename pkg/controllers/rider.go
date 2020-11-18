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
// @Success 200 {object} AlertResponse
// @Router /live/alerts [get]
// @Security ApiKeyAuth
func (controller RiderController) GetRiderAlerts(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var station models.Station
	db.Preload("Info").First(&station, "Name = ?", "Midtown Station") // find product with code l1212

	json.NewEncoder(w).Encode(station)
}
