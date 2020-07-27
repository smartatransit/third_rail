package models

import (
	"time"
)

type StationDetail struct {
	ID          uint       `json:"-" gorm:"primary_key"`
	StationID   uint       `json:"-" gorm:"not null"`
	Description string     `gorm:"not null"`
	Location    string     `gorm:"unique;not null"`
	Distance    float64    `json:",omitempty" gorm:"-"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-" sql:"index"`
}
