package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/smartatransit/third_rail/pkg/models"

	"github.com/smartatransit/third_rail/pkg/transformers"
)

type StaticController struct {
}

// GetStaticScheduleByStation godoc
// @Summary Get Static Schedule By Station
// @Description Get MARTA's scheduled times for arrival for all stations
// @Produce  json
// @Success 200 {object} Response
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
// @Success 200 {object} LinesResponse
// @Router /static/lines [get]
// @Security ApiKeyAuth
func (controller StaticController) GetLines(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var lines []models.Line
	if err := db.Find(&lines).Error; err != nil {
		logrus.Error("Failed to fetch lines:", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			ErrorID: 1,
			Message: "an unknown error occurred",
		})
		return
	}

	response := LinesResponse{linesData{Lines: lines}}

	json.NewEncoder(w).Encode(response)
}

// GetDirections godoc
// @Summary Get Directions
// @Description Get all available directions
// @Produce  json
// @Success 200 {object} DirectionsResponse
// @Router /static/directions [get]
// @Security ApiKeyAuth
func (controller StaticController) GetDirections(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var directions []models.Direction
	if err := db.Find(&directions).Error; err != nil {
		logrus.Error("Failed to fetch directions:", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			ErrorID: 1,
			Message: "an unknown error occurred",
		})
		return
	}

	response := DirectionsResponse{directionsData{Directions: directions}}

	json.NewEncoder(w).Encode(response)
}

// GetStations godoc
// @Summary Get Stations
// @Description Get all available stations
// @Produce  json
// @Success 200 {object} StationsResponse
// @Router /static/stations [get]
// @Security ApiKeyAuth
func (controller StaticController) GetStations(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var stations []models.Station
	if err := db.Preload("Lines").Find(&stations).Error; err != nil {
		logrus.Error("Failed to fetch stations:", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			ErrorID: 1,
			Message: "an unknown error occurred",
		})
		return
	}

	response := StationsResponse{stationsData{Stations: stations}}

	json.NewEncoder(w).Encode(response)
}

// GetLocations implements the
// @Summary Get nearest station given a lat and long, expressed in
// @Description Get nearest station given a lat and lng
// @Produce  json
// @Param latitutde query number true "Latitude"
// @Param longitude query number true "Longitude"
// @Failure 400
// @Success 200 {object} StationsLocationResponse
// @Router /static/stations/location [get]
// @Security ApiKeyAuth
func (controller StaticController) GetLocations(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := req.URL.Query()

	lat, err := strconv.ParseFloat(v.Get("latitude"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			ErrorID: 1,
			Message: "URL parameter `latitude` must be provided and must be a valid decimal number",
		})
		return
	}

	long, err := strconv.ParseFloat(v.Get("longitude"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			ErrorID: 2,
			Message: "URL parameter `longitude` must be provided and must be a valid decimal number",
		})
		return
	}

	var stations []models.Station
	if err := db.Preload("Lines").Preload("Detail").Find(&stations).Error; err != nil {
		logrus.Error("Failed to fetch stations to sort by location:", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			ErrorID: 3,
			Message: "an unknown error occurred",
		})
		return
	}

	sortedStations := transformers.SortStationsByDistance(lat, long, stations)

	response := StationsLocationResponse{sortedStations}

	json.NewEncoder(w).Encode(response)
}
