package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/smartatransit/third_rail/pkg/controllers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStaticController_GetLocations(t *testing.T) {
	app := setUpStaticAPI(t)

	request, _ := http.NewRequest("GET", "/static/stations/location", nil)
	query := request.URL.Query()
	query.Add("latitude", "33.782840")
	query.Add("longitude", "-84.387830")
	request.URL.RawQuery = query.Encode()

	response := httptest.NewRecorder()
	app.Router.ServeHTTP(response, request)

	parsedResponse := controllers.StationsLocationResponse{}

	_ = json.NewDecoder(response.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, 38, len(parsedResponse.Data), "Expected a length of 38")
	assert.Equal(t, "Midtown", parsedResponse.Data[0].Name, "Expected closest station of Midtown Station")
	assert.Equal(t, 707.7872501578765, parsedResponse.Data[0].Detail.Distance, "Expected a distance of 707.78")
}

func TestStaticController_GetStations(t *testing.T) {
	app := setUpStaticAPI(t)

	request, _ := http.NewRequest("GET", "/static/stations", nil)

	response := httptest.NewRecorder()
	app.Router.ServeHTTP(response, request)

	parsedResponse := controllers.StationsResponse{}

	_ = json.NewDecoder(response.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, 38, len(parsedResponse.Data.Stations), "Expected a length of 38")
}

func TestStaticController_GetLines(t *testing.T) {
	app := setUpStaticAPI(t)

	request, _ := http.NewRequest("GET", "/static/lines", nil)

	response := httptest.NewRecorder()
	app.Router.ServeHTTP(response, request)

	parsedResponse := controllers.LinesResponse{}

	_ = json.NewDecoder(response.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, 4, len(parsedResponse.Data.Lines), "Expected a length of 4")
}

func TestStaticController_GetDirections(t *testing.T) {

	app := setUpStaticAPI(t)
	request, _ := http.NewRequest("GET", "/static/directions", nil)
	response := httptest.NewRecorder()
	app.Router.ServeHTTP(response, request)

	parsedResponse := controllers.DirectionsResponse{}

	_ = json.NewDecoder(response.Body).Decode(&parsedResponse)

	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, 4, len(parsedResponse.Data.Directions), "Expected a length of 4")
}

func setUpStaticAPI(t *testing.T) (app *App) {
	db, err := gorm.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatal(err)
	}

	app = &App{DB: db}

	app.Start(true, func() {
		app.Router = mux.NewRouter()
		app.MountStaticRoutes()
	})

	return
}
