package models

import "github.com/jinzhu/gorm"

type StationAlias struct {
	gorm.Model
	StationDetailID   uint `gorm:"not null"`
	Alias       string
	Description string
}
