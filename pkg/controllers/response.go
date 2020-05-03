package controllers

import "github.com/smartatransit/third_rail/pkg/schemas"

type response struct {
	Data []responseData `json:"data"`
}

type responseData struct {
	Schedule schemas.Schedule `json:"schedule"`
	Station  schemas.Station  `json:"station"`
}

type linesResponse struct {
	Data linesData `json:"data"`
}

type linesData struct {
	Lines []string `json:"lines"`
}

type stationsResponse struct {
	Data stationsData `json:"data"`
}

type stationsData struct {
	Stations []string `json:"stations"`
}

type directionsResponse struct {
	Data directionsData `json:"data"`
}

type directionsData struct {
	Directions []string `json:"directions"`
}
