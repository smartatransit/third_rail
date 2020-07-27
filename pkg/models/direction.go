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

type Direction struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	Feedback  []Feedback `json:",omitempty"`
	Lines     []Line     `json:",omitempty" gorm:"many2many:line_directions"`
	Name      string     `gorm:"not null"`
	Aliases   []Alias    `json:",omitempty" gorm:"polymorphic:NamedElement;"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

type Directions []Direction

func (d Directions) String(i int) string {
	return d[i].Name
}

func (d Directions) Len() int {
	return len(d)
}

func FindDirectionByName(name string, db *gorm.DB) (Direction, error) {
	db = db.Set("gorm:auto_preload", true)

	//Try by name
	var directions Directions
	db.Find(&directions)

	for _, s := range directions {
		if strings.ToUpper(s.Name) == strings.ToUpper(name) {
			return s, nil
		}
	}

	//Try by alias
	log.Infof("No Direction name or description match for %s; Searching by alias", name)

	var aliases Aliases

	db.Find(&aliases, "named_element_type = ?", "directions")

	var matches fuzzy.Matches

	results := fuzzy.FindFrom(name, aliases)

	if len(results) > 0 {
		for _, match := range results {
			log.Infof("Direction Name: %s, Match: %s, Score: %d", name, aliases[match.Index].Alias, match.Score)
			matches = append(matches, match)
		}

		var direction Direction
		db.Find(&direction, aliases[matches[0].Index].NamedElementID)

		return direction, nil

	} else {
		err := errors.New(fmt.Sprintf("No station match for name %s", name))
		log.Error(err)
		return Direction{}, err
	}
}
