package models

import "github.com/jinzhu/gorm"

type Line struct {
	gorm.Model
	Feedback   []Feedback  `json:",omitempty"`
	Directions []Direction `gorm:"many2many:line_directions"`
	Stations   []Station   `gorm:"many2man:station_lines"`
	Name       string      `gorm:"not null"`
}

