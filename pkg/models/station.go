package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sahilm/fuzzy"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type Station struct {
	ID        uint          `gorm:"primary_key"`
	Feedback  []Feedback    `json:",omitempty"`
	Detail    StationDetail `json:",omitempty"`
	Aliases   []Alias       `gorm:"polymorphic:NamedElement;"`
	Lines     []Line        `gorm:"many2many:station_lines;not null;"`
	Name      string        `gorm:"unique;not null"`
	CreatedAt time.Time     `json:"-"`
	UpdatedAt time.Time     `json:"-"`
	DeletedAt *time.Time    `json:"-" sql:"index"`
}

type Stations []Station

func (s Stations) String(i int) string {
	return s[i].Name
}

func (s Stations) Len() int {
	return len(s)
}

func FindStationByName(name string, db *gorm.DB) (Station, error) {
	db = db.Set("gorm:auto_preload", true)

	//Try by name or description
	var stations Stations
	db.Find(&stations)

	for _, s := range stations {
		upperName := strings.ToUpper(name)
		if strings.ToUpper(s.Name) == upperName || strings.ToUpper(s.Detail.Description) == upperName ||
			strings.ToUpper(s.Name+" Station") == upperName || strings.ToUpper(s.Detail.Description+" Station") == upperName {
			return s, nil
		}
	}

	//Try by alias
	log.Infof("No name or description match for station %s; Searching by alias", name)

	var aliases Aliases

	db.Find(&aliases, "named_element_type = ?", "stations")
	var matches fuzzy.Matches

	results := fuzzy.FindFrom(name, aliases)

	if len(results) > 0 {
		for _, match := range results {
			log.Infof("Station Name: %s, Match: %s, Score: %d", name, aliases[match.Index].Alias, match.Score)
			matches = append(matches, match)
		}

		var station Station
		db.Find(&station, aliases[matches[0].Index].NamedElementID)

		return station, nil

	} else {
		err := errors.New(fmt.Sprintf("No station match for name %s", name))
		log.Error(err)
		return Station{}, err
	}
}
