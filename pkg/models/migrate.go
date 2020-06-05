package models

import (
	"github.com/jinzhu/gorm"
)

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
