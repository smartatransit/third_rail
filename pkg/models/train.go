package models

import "time"

// Can't use gorm.Model since we want TrainID as the primary key
type Train struct {
	TrainID   uint `gorm:"primary_key:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

