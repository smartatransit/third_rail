package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/gomarta"
	"github.com/smartatransit/third_rail/pkg/clients"
)

type AdminController struct {
	MartaClient clients.MartaClient
	AdminKey    string
}

//HealthResponse represents a response to the health-check endpoint
type HealthResponse struct {
	Statuses []Status `json:"statuses"`
}

//Status represents a single system status
type Status struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Healthy     bool        `json:"healthy"`
	Metadata    interface{} `json:"metadata,omitempty"`
}

func (ac AdminController) checkAdmin(key string) bool {
	return ac.AdminKey == key
}

func (ac AdminController) IngestEvent(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	adminKey := req.Header.Get("key")

	if !ac.checkAdmin(adminKey) {
		w.WriteHeader(http.StatusUnauthorized)
	}

	var event gomarta.Train
	err := json.NewDecoder(req.Body).Decode(event)

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	var resp Response

	json.NewEncoder(w).Encode(resp)
}

//Health responds with a variety of internal statuses
func (ac AdminController) Health(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var statuses []Status
	defer func() {
		if len(statuses) == 0 {
			statuses = append(statuses, Status{
				Name:        "database",
				Description: "postgres backend",
				Healthy:     false,
			})
		}

		ac.writeJSONResponse(w, http.StatusOK, HealthResponse{Statuses: statuses})
	}()

	_, err := ac.MartaClient.GetTrains()
	if err != nil {
		statuses = append(statuses, Status{
			Name:        "marta_client",
			Description: err.Error(),
			Healthy:     false,
		})
	}

	statuses = append(statuses, Status{
		Name:        "database",
		Description: "postgres backend",
		Healthy:     true,
	})
}

func (ac AdminController) writeJSONResponse(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Errorf("failed writing response: %s", err.Error())
		return
	}
}
