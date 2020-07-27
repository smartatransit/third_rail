package controllers

import (
	"github.com/smartatransit/third_rail/pkg/models"
	"github.com/smartatransit/third_rail/pkg/schemas/marta_schemas"
)

type Response struct {
	Data []responseData `json:"data"`
}

type responseData struct {
	Schedule marta_schemas.Schedule `json:"schedule"`
	Station  marta_schemas.Station  `json:"station"`
}

type LinesResponse struct {
	Data linesData `json:"data"`
}

type linesData struct {
	Lines []models.Line `json:"lines"`
}

type ScheduleDetail struct {
	Event    models.ScheduleEvent        `json:"event"`
	RealTime *models.RealTimeEventDetail `json:"real_time_detail,omitempty"`
	Static   *models.StaticEventDetail   `json:"static_event_detail,omitempty"`
}

type StationDetails struct {
	Station  models.Station   `json:"station"`
	Schedule []ScheduleDetail `json:"schedule"`
}

type StationDetailsResponse struct {
	Data StationDetails `json:"data"`
}

type StationsResponse struct {
	Data stationsData `json:"data"`
}

type stationsData struct {
	Stations []models.Station `json:"stations"`
}

type DirectionsResponse struct {
	Data directionsData `json:"data"`
}

type directionsData struct {
	Directions []models.Direction `json:"directions"`
}

type StationsLocationResponse struct {
	Data []models.Station `json:"data"`
}

type parkingData struct {
	Station marta_schemas.Station `json:"station"`
	Status  string                `json:"status"`
}

type ParkingResponse struct {
	Data []parkingData `json:"data"`
}

type AlertResponse struct {
	Data marta_schemas.Alerts `json:"data"`
}
