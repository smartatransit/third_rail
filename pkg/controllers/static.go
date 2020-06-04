package controllers

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/smartatransit/third_rail/pkg/models"
	"net/http"
	"strconv"

	"github.com/smartatransit/third_rail/pkg/transformers"
)

type StaticController struct {
}

// GetStaticScheduleByStation godoc
// @Summary Get Static Schedule By Station
// @Description Get MARTA's scheduled times for arrival for all stations
// @Produce  json
// @Success 200 {object} response
// @Router /static/schedule/station [get]
// @Security ApiKeyAuth
func (controller StaticController) GetStaticScheduleByStation(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := req.URL.Query()

	schedule := v.Get("schedule")
	stationName := v.Get("station_name")

	if len(schedule) < 1 || len(stationName) < 1 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	json.NewEncoder(w).Encode(Response{})
}

// GetLines godoc
// @Summary Get Lines
// @Description Get all available lines
// @Produce  json
// @Success 200 {object} linesResponse
// @Router /static/lines [get]
// @Security ApiKeyAuth
func (controller StaticController) GetLines(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var lines []models.Line
	db.Find(&lines)
	response := LinesResponse{linesData{Lines: lines}}

	json.NewEncoder(w).Encode(response)
}

// GetDirections godoc
// @Summary Get Directions
// @Description Get all available directions
// @Produce  json
// @Success 200 {object} directionsResponse
// @Router /static/directions [get]
// @Security ApiKeyAuth
func (controller StaticController) GetDirections(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var directions []models.Direction
	db.Find(&directions)

	response := DirectionsResponse{directionsData{Directions: directions}}

	json.NewEncoder(w).Encode(response)
}

// GetStations godoc
// @Summary Get Stations
// @Description Get all available stations
// @Produce  json
// @Success 200 {object} stationsResponse
// @Router /static/stations [get]
// @Security ApiKeyAuth
func (controller StaticController) GetStations(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var stations []models.Station
	//db.Preload("Lines").Find(&stations)
	db.Find(&stations)

	response := StationsResponse{stationsData{Stations: stations}}

	json.NewEncoder(w).Encode(response)
}

// GetLocations godoc
// @Summary Get nearest station given a lat and lng
// @Description Get nearest station given a lat and lng
// @Produce  json
// @Param latitutde query int true "Latitude"
// @Param longitude query int true "Longitude"
// @Failure 400
// @Success 200 {object} stationsLocationResponse
// @Router /static/location [get]
// @Security ApiKeyAuth
func (controller StaticController) GetLocations(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := req.URL.Query()

	lat, latErr := strconv.ParseFloat(v.Get("latitude"), 64)
	long, longErr := strconv.ParseFloat(v.Get("longitude"), 64)

	//mev := validators.NewMartaEntitiesValidator()

	if longErr != nil || latErr != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	lt := transformers.NewLocationTransformer()

	response := StationsLocationResponse{lt.GetNearestLocations(lat, long)}

	json.NewEncoder(w).Encode(response)
}
