package models

import "github.com/jinzhu/gorm"

type Station struct {
	gorm.Model
	Feedback []Feedback
	Detail   StationDetail
	Lines    []Line `gorm:"many2many:station_lines;not null"`
	Name     string `gorm:"unique;not null"`
}

