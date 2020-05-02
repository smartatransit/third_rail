package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/smartatransit/gomarta"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/clients/marta_client"
	"github.com/smartatransit/third_rail/pkg/schemas/marta_schemas"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLiveController_GetScheduleByLine(t *testing.T) {
	harness := marta_client.MartaAPIHarness{
		GetTrains: func() (r gomarta.TrainAPIResponse, e error) {
			r = gomarta.TrainAPIResponse{
				gomarta.Train{
					Destination:    "North Springs Station",
					Direction:      "N",
					EventTime:      "5/2/2020 1:01:01 AM",
					Line:           "Red",
					NextArrival:    "01:05:17 AM",
					Station:        "Medical Center Station",
					TrainID:        "404306",
					WaitingSeconds: "236",
					WaitingTime:    "3 min",
				},
				gomarta.Train{
					Destination:    "Doraville Station",
					Direction:      "N",
					EventTime:      "5/2/2020 1:01:01 AM",
					Line:           "Gold",
					NextArrival:    "01:05:17 AM",
					Station:        "Midtown Station",
					TrainID:        "404307",
					WaitingSeconds: "236",
					WaitingTime:    "3 min",
				},
			}
			return
		},
		GetAlerts: func() (a marta_schemas.Alerts, e error) {
			return
		},
	}

	martaClient := marta_client.CreateMartaAPITestClient(harness)

	request, _ := http.NewRequest("GET", "/live/schedule/line/red", nil)
	resp := processLiveRequest(request, martaClient)

	parsedResponse := response{}

	_ = json.NewDecoder(resp.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, resp.Code, "OK response is expected")
	assert.Equal(t, 1, len(parsedResponse.Data), "Expected a length of 1")
	assert.Equal(t, "Red", parsedResponse.Data[0].Station.Line, "Expected a line of 'Red'")
}

func TestLiveController_GetScheduleByStation(t *testing.T) {
	harness := marta_client.MartaAPIHarness{
		GetTrains: func() (r gomarta.TrainAPIResponse, e error) {
			r = gomarta.TrainAPIResponse{
				gomarta.Train{
					Destination:    "North Springs Station",
					Direction:      "N",
					EventTime:      "5/2/2020 1:01:01 AM",
					Line:           "Red",
					NextArrival:    "01:05:17 AM",
					Station:        "Medical Center Station",
					TrainID:        "404306",
					WaitingSeconds: "236",
					WaitingTime:    "3 min",
				},
				gomarta.Train{
					Destination:    "Doraville Station",
					Direction:      "N",
					EventTime:      "5/2/2020 1:01:01 AM",
					Line:           "Gold",
					NextArrival:    "01:05:17 AM",
					Station:        "Midtown Station",
					TrainID:        "404307",
					WaitingSeconds: "236",
					WaitingTime:    "3 min",
				},
			}
			return
		},
		GetAlerts: func() (a marta_schemas.Alerts, e error) {
			return
		},
	}

	martaClient := marta_client.CreateMartaAPITestClient(harness)

	request, _ := http.NewRequest("GET", "/live/schedule/station/Medical Center Station", nil)
	resp := processLiveRequest(request, martaClient)

	parsedResponse := response{}

	_ = json.NewDecoder(resp.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, resp.Code, "OK response is expected")
	assert.Equal(t, 1, len(parsedResponse.Data), "Expected a station length of 1")
	assert.Equal(t, "Medical Center Station", parsedResponse.Data[0].Station.Name, "Expected a station name of 'Medical Center Station")
}

func TestLiveController_GetAlerts(t *testing.T) {
	harness := marta_client.MartaAPIHarness{
		GetTrains: func() (r gomarta.TrainAPIResponse, e error) {
			return
		},
		GetAlerts: func() (a marta_schemas.Alerts, e error) {
			a.Bus = []marta_schemas.BusAlert{
				{
					ID:      "135400",
					Title:   "5/2/2020 - Service Alert",
					Desc:    "ROUTE 2: WESTBOUND TO NORTH AVENUE STATION @ 0930 WILL BE DELAYED.",
					Expires: "2020-05-02T10:10:58-07:00",
				},
			}

			a.Rail = []marta_schemas.RailAlert{
				{
					ID:      "135392",
					Title:   "5/2/2020 - Service Alert",
					Desc:    "Park & Ride available at 23 MARTA rail stations. For a list of free & long-term locations, visit http://ow.ly/DGMp30ashGr",
					Expires: "2020-05-03T06:05:00-07:00",
				},
			}

			return
		},
	}

	martaClient := marta_client.CreateMartaAPITestClient(harness)

	request, _ := http.NewRequest("GET", "/live/alerts", nil)
	resp := processLiveRequest(request, martaClient)

	parsedResponse := alertResponse{}

	_ = json.NewDecoder(resp.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, resp.Code, "OK response is expected")
	assert.Equal(t, 1, len(parsedResponse.Data.Rail), "Expected a rail alert length of 1")
	assert.Equal(t, 1, len(parsedResponse.Data.Bus), "Expected a bus alert length of 1")
}

func processLiveRequest(request *http.Request, mc clients.MartaClient) (response *httptest.ResponseRecorder) {
	router := mux.NewRouter()

	liveController := LiveController{MartaClient: mc}
	liveRouter := router.PathPrefix("/live").Subrouter()
	liveRouter.HandleFunc("/schedule/line/{line}", liveController.GetScheduleByLine).Methods("GET")
	liveRouter.HandleFunc("/schedule/station/{station}", liveController.GetScheduleByStation).Methods("GET")
	liveRouter.HandleFunc("/alerts", liveController.GetAlerts).Methods("GET")

	response = httptest.NewRecorder()
	router.ServeHTTP(response, request)

	return
}
