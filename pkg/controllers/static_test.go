package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStaticController_GetLocations(t *testing.T) {

	request, _ := http.NewRequest("GET", "/static/stations/location", nil)
	query := request.URL.Query()
	query.Add("latitude", "33.782840")
	query.Add("longitude", "-84.387830")
	request.URL.RawQuery = query.Encode()

	resp := processStaticRequest(request)

	parsedResponse := stationsLocationResponse{}

	_ = json.NewDecoder(resp.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, resp.Code, "OK response is expected")
	assert.Equal(t, 38, len(parsedResponse.Data), "Expected a length of 38")
	assert.Equal(t, "Midtown Station", parsedResponse.Data[0].StationName, "Expected closest station of Midtown Station")
	assert.Equal(t, 707.7872501578765, parsedResponse.Data[0].Distance, "Expected a distance of 707.78")
}

func TestStaticController_GetStations(t *testing.T) {

	request, _ := http.NewRequest("GET", "/static/stations", nil)

	resp := processStaticRequest(request)

	parsedResponse := stationsResponse{}

	_ = json.NewDecoder(resp.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, resp.Code, "OK response is expected")
	assert.Equal(t, 38, len(parsedResponse.Data.Stations), "Expected a length of 38")
}

func TestStaticController_GetLines(t *testing.T) {

	request, _ := http.NewRequest("GET", "/static/lines", nil)

	resp := processStaticRequest(request)

	parsedResponse := linesResponse{}

	_ = json.NewDecoder(resp.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, resp.Code, "OK response is expected")
	assert.Equal(t, 4, len(parsedResponse.Data.Lines), "Expected a length of 4")
}

func TestStaticController_GetDirections(t *testing.T) {

	request, _ := http.NewRequest("GET", "/static/directions", nil)

	resp := processStaticRequest(request)

	parsedResponse := directionsResponse{}

	_ = json.NewDecoder(resp.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, resp.Code, "OK response is expected")
	assert.Equal(t, 4, len(parsedResponse.Data.Directions), "Expected a length of 4")
}

func processStaticRequest(request *http.Request) (response *httptest.ResponseRecorder) {
	router := mux.NewRouter()

	staticController := StaticController{}
	staticRouter := router.PathPrefix("/static").Subrouter()
	staticRouter.HandleFunc("/schedule/station", staticController.GetStaticScheduleByStation).Methods("GET")
	staticRouter.HandleFunc("/lines", staticController.GetLines).Methods("GET")
	staticRouter.HandleFunc("/directions", staticController.GetDirections).Methods("GET")
	staticRouter.HandleFunc("/stations", staticController.GetStations).Methods("GET")
	staticRouter.HandleFunc("/stations/location", staticController.GetLocations).Methods("GET")

	response = httptest.NewRecorder()
	router.ServeHTTP(response, request)

	return
}
