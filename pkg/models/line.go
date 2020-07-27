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

type Line struct {
	ID         uint        `json:"-" gorm:"primary_key"`
	Feedback   []Feedback  `json:",omitempty"`
	Directions []Direction `json:",omitempty" gorm:"many2many:line_directions"`
	Stations   []Station   `json:",omitempty" gorm:"many2man:station_lines"`
	Name       string      `gorm:"not null"`
	Aliases    []Alias     `json:",omitempty" gorm:"polymorphic:NamedElement;"`
	CreatedAt  time.Time   `json:"-"`
	UpdatedAt  time.Time   `json:"-"`
	DeletedAt  *time.Time  `json:"-" sql:"index"`
}

type Lines []Line

func (l Lines) String(i int) string {
	return l[i].Name
}

func (l Lines) Len() int {
	return len(l)
}

func FindLineByName(name string, db *gorm.DB) (Line, error) {
	db = db.Set("gorm:auto_preload", true)

	//Try by name
	var lines Lines
	db.Find(&lines)

	for _, l := range lines {
		if strings.ToUpper(l.Name) == strings.ToUpper(name) {
			return l, nil
		}
	}

	//Try by alias
	log.Infof("No Line name or description match for %s; Searching by alias", name)

	var aliases Aliases

	db.Find(&aliases, "named_element_type = ?", "lines")

	var matches fuzzy.Matches

	results := fuzzy.FindFrom(name, aliases)

	if len(results) > 0 {
		for _, match := range results {
			log.Infof("Line Name: %s, Match: %s, Score: %d", name, aliases[match.Index].Alias, match.Score)
			matches = append(matches, match)
		}

		var line Line
		db.Find(&line, aliases[matches[0].Index].NamedElementID)

		return line, nil

	} else {
		err := errors.New(fmt.Sprintf("No station match for name %s", name))
		log.Error(err)
		return Line{}, err
	}
}
