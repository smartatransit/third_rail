package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/third_rail/docs"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/clients/marta_client"
	"github.com/smartatransit/third_rail/pkg/clients/twitter_client"
	"github.com/smartatransit/third_rail/pkg/controllers"
	"github.com/smartatransit/third_rail/pkg/daemons"
	"github.com/smartatransit/third_rail/pkg/middleware"
	"github.com/smartatransit/third_rail/pkg/models"
	"github.com/smartatransit/third_rail/pkg/seed"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	Router        *mux.Router
	DB            *gorm.DB
	MartaClient   clients.MartaClient
	TwitterClient clients.TwitterClient
	Options       Options
}

// @title SMARTA API
// @version 1.0
// @description API to serve you SMARTA data

// @contact.name SMARTA Support
// @contact.email smartatransit@gmail.com

// @license.name GNU General Public License v3.0
// @license.url https://github.com/smartatransit/third_rail/blob/master/LICENSE

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func (app *App) Start(customRouter func()) {
	// We do this here rather than in Initialize so that it only happens
	// when we start the app. Forcing preloading while seeding or migrating
	// could slow things down.
	app.DB.Set("gorm:auto_preload", true)

	// Most of the Swagger config lives in hard-coded comments, but since
	// the hostname is variable, we configure it at runtime:
	docs.SwaggerInfo.Host = app.Options.Servicedomain

	if customRouter != nil {
		customRouter()
	} else {
		if app.MartaClient == nil {
			app.MartaClient = marta_client.GetMartaClient(app.Options.MartaAPIKey, app.Options.MartaCacheTTL)
		}

		if app.TwitterClient == nil {
			app.TwitterClient = twitter_client.GetTwitterClient(app.Options.TwitterClientID, app.Options.TwitterClientSecret, app.Options.TwitterCacheTTL)
		}

		app.Router = mux.NewRouter()
		app.mountLiveRoutes()
		app.MountStaticRoutes()
		app.mountSmartRoutes()
		app.mountRiderRoutes()
		app.mountSwaggerRoutes()
		app.mountAdminRoutes()

		if app.Options.RailRunner {
			rr := daemons.RailRunner{MartaClient: app.MartaClient, DB: app.DB}
			rr.Start()
		}

		app.serve()
	}
}

func (app *App) Initialize() error {
	log.SetFormatter(&log.JSONFormatter{})

	if app.DB == nil {
		var err error
		app.DB, err = gorm.Open("postgres", app.Options.DBConnectionString)
		if err != nil {
			return fmt.Errorf("Could not connect database: %w", err)
		}
	}

	return nil
}

func (app *App) InitializeSchema(quiet ...bool) error {
	q := len(quiet) > 0 && quiet[0]

	app.DB.LogMode(!q)
	lg := log.New()
	if q {
		lg.SetLevel(log.ErrorLevel)
	} else {
		lg.SetLevel(log.InfoLevel)
	}

	if err := models.DBMigrate(app.DB, lg); err != nil {
		return fmt.Errorf("failed migrating database: %w", err)
	}
	if err := seed.Seed(app.DB, lg); err != nil {
		return fmt.Errorf("failed seeding database: %w", err)
	}

	lg.Info("Success!")
	return nil
}

func (app *App) mountLiveRoutes() {
	if app.MartaClient == nil {
		log.Fatal("No MARTA client present - unable to mount Live routes.")
	}

	liveController := controllers.LiveController{MartaClient: app.MartaClient}
	liveRouter := app.Router.PathPrefix("/live").Subrouter()
	liveRouter.HandleFunc("/schedule/line/{line}", liveController.GetScheduleByLine).Methods("GET")
	liveRouter.HandleFunc("/schedule/station/{station}", liveController.GetScheduleByStation).Methods("GET")
	liveRouter.HandleFunc("/alerts", liveController.GetAlerts).Methods("GET")
}

func (app *App) MountStaticRoutes() {
	staticController := controllers.StaticController{}
	staticRouter := app.Router.PathPrefix("/static").Subrouter()
	staticRouter.HandleFunc("/schedule/station", staticController.GetStaticScheduleByStation).Methods("GET")
	staticRouter.HandleFunc("/lines", func(w http.ResponseWriter, r *http.Request) {
		staticController.GetLines(app.DB, w, r)
	}).Methods("GET")
	staticRouter.HandleFunc("/directions", func(w http.ResponseWriter, r *http.Request) {
		staticController.GetDirections(app.DB, w, r)
	}).Methods("GET")
	staticRouter.HandleFunc("/stations", func(w http.ResponseWriter, r *http.Request) {
		staticController.GetStations(app.DB, w, r)
	}).Methods("GET")
	staticRouter.HandleFunc("/stations/location", func(w http.ResponseWriter, r *http.Request) {
		staticController.GetLocations(app.DB, w, r)
	}).Methods("GET")
}

func (app *App) mountSmartRoutes() {
	if app.TwitterClient == nil {
		log.Fatal("No Twitter client present - unable to mount Smart routes.")
	}

	smartController := controllers.SmartController{TwitterClient: app.TwitterClient, MartaClient: app.MartaClient}
	smartRouter := app.Router.PathPrefix("/smart").Subrouter()
	smartRouter.HandleFunc("/parking", func(w http.ResponseWriter, r *http.Request) {
		smartController.GetParkingStatus(app.DB, w, r)
	}).Methods("GET")
	smartRouter.HandleFunc("/station/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		smartController.GetStationDetails(app.DB, w, r)
	}).Methods("GET")
	smartRouter.HandleFunc("/emergencies", smartController.GetEmergencyStatus).Methods("GET")
}

func (app *App) mountRiderRoutes() {
	riderController := controllers.RiderController{}
	riderRouter := app.Router.PathPrefix("/rider").Subrouter()
	riderRouter.HandleFunc("/alerts", func(w http.ResponseWriter, r *http.Request) {
		riderController.GetRiderAlerts(app.DB, w, r)
	}).Methods("GET")
}

func (app *App) mountAdminRoutes() {

	adminController := controllers.AdminController{MartaClient: app.MartaClient, AdminKey: app.Options.AdminAPIKey}
	adminRouter := app.Router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/event/ingest", func(w http.ResponseWriter, r *http.Request) {
		adminController.IngestEvent(app.DB, w, r)
	}).Methods("POST")
}

func (app *App) mountSwaggerRoutes() {
	app.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

func (app *App) serve() {
	fmt.Println("started on port :5000")

	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)(middleware.SuffixMiddleware(app.Router))))
}
