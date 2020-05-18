package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Direction struct {
	//gorm.Model
	DirectionID uint `gorm:"primary_key:true"`
	//Lines []Line `gorm:"many2many:line_directions;association_foreignkey:ID;foreignkey:ID"`
	Name string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Line struct {
	//gorm.Model
	LineID uint `gorm:"primary_key:true"`
	Directions []Direction `gorm:"many2many:line_directions;association_foreignkey:DirectionID;foreignkey:LineID"`
	Name      string      `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Station struct {
	gorm.Model
	Name string        `gorm:"unique;not null"`
	Info StationDetail `gorm:"foreignkey:ID"`
}

type StationDetail struct {
	gorm.Model
	Lines       []Line `gorm:"many2many:station_detail_lines;not null"`
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

type StationFeedback struct {
	gorm.Model
	Source      FeedbackSource `gorm:"foreignkey:ID;not null"`
	Type        FeedbackType   `gorm:"foreignkey:ID;not null"`
	Description string         `gorm:"not null"`
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
	EventType   ScheduleEventSource `gorm:"foreignkey:ID"`
	Destination Station             `gorm:"foreignkey:ID;not null"`
	NextArrival time.Time
	NextStation Station `gorm:"foreignkey:ID;not null"`
}

type RealTimeEventDetail struct {
	gorm.Model
	ScheduleEvent  ScheduleEvent `gorm:"foreignkey:ID;not null"`
	EventTime      time.Time
	Train          Train  `gorm:"foreignkey:TrainID"`
	WaitingSeconds int    `gorm:"not null"`
	WaitingTime    string `gorm:"not null"`
}

type StaticEventDetail struct {
	gorm.Model
	ScheduleEvent      ScheduleEvent `gorm:"foreignkey:ID;not null"`
	ScheduledTime      time.Time
	StaticScheduleType string `gorm:"not null"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Direction{})
	db.AutoMigrate(&Line{})
	//db.Model("line_directions").AddForeignKey("line_id", "lines(id)", "RESTRICT", "RESTRICT")
	//db.Model("line_directions").AddForeignKey("direction_id", "directions(id)", "RESTRICT", "RESTRICT")

	db.AutoMigrate(&Station{})
	db.AutoMigrate(&StationDetail{})
	db.AutoMigrate(&FeedbackType{})
	db.AutoMigrate(&FeedbackSource{})
	db.AutoMigrate(&StationFeedback{})
	db.AutoMigrate(&Train{})
	db.AutoMigrate(&ScheduleEventSource{})
	db.AutoMigrate(&ScheduleEvent{})
	db.AutoMigrate(&RealTimeEventDetail{})
	db.AutoMigrate(&StaticEventDetail{})
	return db
}
