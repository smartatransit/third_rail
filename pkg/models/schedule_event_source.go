package models

import "github.com/jinzhu/gorm"

type ScheduleEventSource struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}
