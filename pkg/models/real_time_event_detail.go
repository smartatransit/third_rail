package models

import (
	"time"
)

type RealTimeEventDetail struct {
	ID              uint          `json:"-" gorm:"primary_key"`
	ScheduleEventID uint          `json:"-"`
	ScheduleEvent   ScheduleEvent `json:"-"`
	EventTime       time.Time
	TrainID         uint `json:"-"`
	Train           Train
	WaitingSeconds  int    `gorm:"not null"`
	WaitingTime     string `gorm:"not null"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
