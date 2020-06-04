package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Direction struct {
	gorm.Model
	Feedback []Feedback
	Lines    []Line `gorm:"many2many:line_directions"`
	Name     string `gorm:"not null"`
}

type Line struct {
	gorm.Model
	Feedback   []Feedback  `json:",omitempty"`
	Directions []Direction `gorm:"many2many:line_directions"`
	Stations   []Station   `gorm:"many2man:station_lines"`
	Name       string      `gorm:"not null"`
}

type Station struct {
	gorm.Model
	Feedback []Feedback
	Detail   StationDetail
	Lines    []Line `gorm:"many2many:station_lines;not null"`
	Name     string `gorm:"unique;not null"`
}

type StationDetail struct {
	gorm.Model
	StationID   uint   `gorm:"not null"`
	Description string `gorm:"not null"`
	Location    string `gorm:"unique;not null"`
}

type FeedbackType struct {
	gorm.Model
	Description string `gorm:"not null"`
}

type FeedbackSource struct {
	gorm.Model
	SourceID   string
	SourceType string
}

type Feedback struct {
	gorm.Model
	StationID   int
	LineID      int
	DirectionID int
	SourceID    int
	Source      FeedbackSource
	TypeID      int
	Type        FeedbackType
	Description string `gorm:"not null"`
	ThumbsUp    int
	ThumbsDown  int
	ExpiresAt   time.Time
}

// Can't use gorm.Model since we want TrainID as the primary key
type Train struct {
	TrainID   uint `gorm:"primary_key:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ScheduleEventSource struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

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

type StaticEventDetail struct {
	gorm.Model
	ScheduleEventID    int
	ScheduleEvent      ScheduleEvent
	ScheduledTime      time.Time
	StaticScheduleType string `gorm:"not null"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.LogMode(false)

	db.AutoMigrate(&Direction{})
	db.AutoMigrate(&Line{})
	db.Table("line_directions").AddForeignKey("line_id", "lines(id)", "RESTRICT", "RESTRICT")
	db.Table("line_directions").AddForeignKey("direction_id", "directions(id)", "RESTRICT", "RESTRICT")

	db.AutoMigrate(&Station{})
	db.AutoMigrate(&StationDetail{})
	db.Model(&StationDetail{}).AddForeignKey("station_id", "stations(id)", "RESTRICT", "RESTRICT")
	db.Table("station_lines").AddForeignKey("line_id", "lines(id)", "RESTRICT", "RESTRICT")
	db.Table("station_lines").AddForeignKey("station_id", "stations(id)", "RESTRICT", "RESTRICT")

	db.AutoMigrate(&FeedbackType{})
	db.AutoMigrate(&FeedbackSource{})
	db.AutoMigrate(&Feedback{})
	db.Model(&Feedback{}).AddForeignKey("source_id", "feedback_sources(id)", "RESTRICT", "RESTRICT")
	db.Model(&Feedback{}).AddForeignKey("type_id", "feedback_types(id)", "RESTRICT", "RESTRICT")
	db.Model(&Feedback{}).AddForeignKey("station_id", "stations(id)", "RESTRICT", "RESTRICT")
	db.Model(&Feedback{}).AddForeignKey("line_id", "lines(id)", "RESTRICT", "RESTRICT")
	db.Model(&Feedback{}).AddForeignKey("direction_id", "directions(id)", "RESTRICT", "RESTRICT")

	db.AutoMigrate(&Train{})

	db.AutoMigrate(&ScheduleEventSource{})
	db.AutoMigrate(&ScheduleEvent{})
	db.Model(&ScheduleEvent{}).AddForeignKey("event_type_id", "schedule_event_sources(id)", "RESTRICT", "RESTRICT")
	db.Model(&ScheduleEvent{}).AddForeignKey("destination_id", "stations(id)", "RESTRICT", "RESTRICT")
	db.Model(&ScheduleEvent{}).AddForeignKey("next_station_id", "stations(id)", "RESTRICT", "RESTRICT")

	db.AutoMigrate(&RealTimeEventDetail{})
	db.AutoMigrate(&RealTimeEventDetail{}).AddForeignKey("schedule_event_id", "schedule_events(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&RealTimeEventDetail{}).AddForeignKey("train_id", "trains(train_id)", "RESTRICT", "RESTRICT")

	db.AutoMigrate(&StaticEventDetail{})
	db.AutoMigrate(&StaticEventDetail{}).AddForeignKey("schedule_event_id", "schedule_events(id)", "RESTRICT", "RESTRICT")

	return db.Set("gorm:auto_preload", true)
}
