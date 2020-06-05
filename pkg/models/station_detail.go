package models

import "github.com/jinzhu/gorm"

type StationDetail struct {
	gorm.Model
	Aliases     []StationAlias
	StationID   uint    `gorm:"not null"`
	Description string  `gorm:"not null"`
	Location    string  `gorm:"unique;not null"`
	Distance    float64 `gorm:"-"`
}
