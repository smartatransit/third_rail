package controllers

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/gomarta"
	"github.com/smartatransit/third_rail/pkg/clients"
	"net/http"
)

type AdminController struct {
	MartaClient clients.MartaClient
	AdminKey    string
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
