package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ScheduleEvent struct {
	gorm.Model
	EventTypeID   int
	EventType     ScheduleEventSource
	DestinationID int
	Destination   Station
	NextArrival   time.Time
	NextStationID int
	NextStation   Station
}

