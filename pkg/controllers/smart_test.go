package controllers

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/mux"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/clients/twitter_client"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSmartController_GetEmergencyStatus(t *testing.T) {
	harness := twitter_client.TwitterHarness{
		Search: func(s string, p *twitter.SearchTweetParams) (r *twitter.Search, e error) {
			assert.Contains(t, p.Query, "from:@martapolice", "Expected target of @martapolice.")

			r = &twitter.Search{
				Statuses: []twitter.Tweet{
					{
						FullText: "Test emergency from MPD",
					},
				},
			}

			return
		},
	}

	twitterClient := twitter_client.CreateTwitterTestClient(harness)

	request, _ := http.NewRequest("GET", "/smart/emergencies", nil)
	resp := processSmartRequest(request, twitterClient)

	parsedResponse := parkingResponse{}

	_ = json.NewDecoder(resp.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, resp.Code, "OK response is expected")
	assert.Equal(t, 1, len(parsedResponse.Data), "Expected a length of 1")
}

func TestSmartController_GetParkingStatus(t *testing.T) {
	harness := twitter_client.TwitterHarness{
		Search: func(s string, p *twitter.SearchTweetParams) (r *twitter.Search, e error) {
			assert.Contains(t, p.Query, "#parkingupdate", "Expected '#parkingupdate' term in query.")
			assert.Contains(t, p.Query, "from:@martaservice", "Expected target of @martaservice.")

			r = &twitter.Search{
				Statuses: []twitter.Tweet{
					{
						FullText: "Test parking update from MARTA",
					},
				},
			}

			return
		},
	}

	twitterClient := twitter_client.CreateTwitterTestClient(harness)

	request, _ := http.NewRequest("GET", "/smart/parking", nil)
	resp := processSmartRequest(request, twitterClient)

	parsedResponse := parkingResponse{}

	_ = json.NewDecoder(resp.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, resp.Code, "OK response is expected")
	assert.Equal(t, 1, len(parsedResponse.Data), "Expected a length of 1")
}

func processSmartRequest(request *http.Request, tc clients.TwitterClient) (response *httptest.ResponseRecorder) {
	router := mux.NewRouter()

	smartController := SmartController{TwitterClient: tc}
	smartRouter := router.PathPrefix("/smart").Subrouter()
	smartRouter.HandleFunc("/parking", smartController.GetParkingStatus).Methods("GET")
	smartRouter.HandleFunc("/emergencies", smartController.GetEmergencyStatus).Methods("GET")

	response = httptest.NewRecorder()
	router.ServeHTTP(response, request)

	return
}
