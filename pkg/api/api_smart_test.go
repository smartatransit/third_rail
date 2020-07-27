package api

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/clients/twitter_client"
	"github.com/smartatransit/third_rail/pkg/controllers"
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

	app := setUpSmartAPI(t, twitterClient)

	request, _ := http.NewRequest("GET", "/smart/emergencies", nil)
	response := httptest.NewRecorder()
	app.Router.ServeHTTP(response, request)

	parsedResponse := controllers.ParkingResponse{}

	_ = json.NewDecoder(response.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, response.Code, "OK response is expected")
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

	app := setUpSmartAPI(t, twitterClient)

	request, _ := http.NewRequest("GET", "/smart/parking", nil)
	response := httptest.NewRecorder()
	app.Router.ServeHTTP(response, request)

	parsedResponse := controllers.ParkingResponse{}

	_ = json.NewDecoder(response.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, 1, len(parsedResponse.Data), "Expected a length of 1")
}

func setUpSmartAPI(t *testing.T, twitterClient clients.TwitterClient) (app *App) {
	db, err := gorm.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatal(err)
	}

	app = &App{DB: db, TwitterClient: twitterClient}

	app.Start(true, func() {
		app.Router = mux.NewRouter()
		app.mountSmartRoutes()
	})

	return
}
