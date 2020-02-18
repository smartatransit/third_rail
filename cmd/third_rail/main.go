package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/smartatransit/gomarta"
	"github.com/smartatransit/third_rail/pkg/controllers"
	"github.com/smartatransit/third_rail/pkg/middleware"
	"log"
	"net/http"
	"os"
)

type MartaClient interface{}

var martaClient MartaClient

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if martaClient == nil {
		martaClient = getMartaClient()
	}

	mountAndServe(martaClient)
}

func mountAndServe(client MartaClient) {
	router := mux.NewRouter()

	liveRouter := router.PathPrefix("/live").Subrouter()
	liveRouter.HandleFunc("/schedule/line/{line}", controllers.GetScheduleByLine).Methods("GET")
	liveRouter.HandleFunc("/schedule/station/{station}", controllers.GetScheduleByStation).Methods("GET")

	staticRouter := router.PathPrefix("/static").Subrouter()
	staticRouter.HandleFunc("/schedule/station", controllers.GetStaticScheduleByStation).Methods("GET")
	staticRouter.HandleFunc("/lines", controllers.GetLines).Methods("GET")
	staticRouter.HandleFunc("/directions", controllers.GetDirections).Methods("GET")
	staticRouter.HandleFunc("/stations", controllers.GetStations).Methods("GET")
	staticRouter.HandleFunc("/stations/location", controllers.GetLocations).Methods("GET")

	fmt.Println("started on port :5000")

	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)(middleware.SuffixMiddleware(router))))
}

func getMartaClient() MartaClient {
	return gomarta.NewDefaultClient(os.Getenv("MARTA_API_KEY"))
}

