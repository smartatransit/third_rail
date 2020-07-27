package models

import "time"

type ScheduleEventSource struct {
	ID          uint       `json:"-" gorm:"primary_key"`
	Name        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-" sql:"index"`
}
