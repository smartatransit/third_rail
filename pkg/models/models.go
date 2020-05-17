package models

import "github.com/jinzhu/gorm"

type Station struct {
	gorm.Model
	Name string      `gorm:"unique;not null"`
	Info StationInfo `gorm:"foreignkey:ID"`
}

type StationInfo struct {
	gorm.Model
	Description string `gorm:"not null"`
	Location    string `gorm:"unique;not null"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Station{})
	db.AutoMigrate(&StationInfo{})
	return db
}
