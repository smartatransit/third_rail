package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RealTimeEventDetail struct {
	gorm.Model
	ScheduleEventID int
	ScheduleEvent   ScheduleEvent
	EventTime       time.Time
	TrainID         int
	Train           Train
	WaitingSeconds  int    `gorm:"not null"`
	WaitingTime     string `gorm:"not null"`
}
