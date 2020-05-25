package controllers

import (
	"github.com/smartatransit/third_rail/pkg/models"
	"github.com/smartatransit/third_rail/pkg/schemas/marta_schemas"
)

type response struct {
	Data []responseData `json:"data"`
}

type responseData struct {
	Schedule marta_schemas.Schedule `json:"schedule"`
	Station  marta_schemas.Station  `json:"station"`
}

type linesResponse struct {
	Data linesData `json:"data"`
}

type linesData struct {
	Lines []models.Line `json:"lines"`
}

type stationsResponse struct {
	Data stationsData `json:"data"`
}

type stationsData struct {
	Stations []models.Station `json:"stations"`
}

type directionsResponse struct {
	Data directionsData `json:"data"`
}

type directionsData struct {
	Directions []models.Direction `json:"directions"`
}

type stationsLocationResponse struct {
	Data []marta_schemas.StationLocation `json:"data"`
}

type parkingData struct {
	Station marta_schemas.Station `json:"station"`
	Status  string                `json:"status"`
}

type parkingResponse struct {
	Data []parkingData `json:"data"`
}

type alertResponse struct {
	Data marta_schemas.Alerts `json:"data"`
}
