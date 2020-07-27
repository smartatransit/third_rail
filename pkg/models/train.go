package models

import "time"

type Train struct {
	ID        uint       `gorm:"primary_key:true"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
