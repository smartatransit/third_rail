package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type StaticEventDetail struct {
	gorm.Model
	ScheduleEventID    int
	ScheduleEvent      ScheduleEvent
	ScheduledTime      time.Time
	StaticScheduleType string `gorm:"not null"`
}
