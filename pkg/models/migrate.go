package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func DBMigrate(db *gorm.DB, log *log.Logger) error {
	log.Info("Creating tables and indexes...")

	var err error
	if err = db.AutoMigrate(&Direction{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `Direction`: %w", err)
	}
	if err = db.AutoMigrate(&Line{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `Line`: %w", err)
	}

	ldTable := db.Table("line_directions")
	if err = ldTable.AddForeignKey("line_id", "lines(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("adding line foreign key: %w", err)
	}
	if err = ldTable.AddForeignKey("direction_id", "directions(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("adding direction foreign key: %w", err)
	}

	if err = db.AutoMigrate(&Station{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `Station`: %w", err)
	}
	if err = db.AutoMigrate(&StationDetail{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `StationDetail`: %w", err)
	}
	if err = db.Model(&StationDetail{}).AddForeignKey("station_id", "stations(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding station/station-detatil foreign key: %w", err)
	}
	slTable := db.Table("station_lines")
	if err = slTable.AddForeignKey("line_id", "lines(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding line-station foreign key: %w", err)
	}
	if err = slTable.AddForeignKey("station_id", "stations(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding station-line foreign key: %w", err)
	}

	if err = db.AutoMigrate(&Alias{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `Alias`: %w", err)
	}
	if err = db.AutoMigrate(&FeedbackType{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `FeedbackType`: %w", err)
	}
	if err = db.AutoMigrate(&FeedbackSource{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `FeedbackSource`: %w", err)
	}
	if err = db.AutoMigrate(&Feedback{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `Feedback`: %w", err)
	}

	fbModel := db.Model(&Feedback{})
	if err = fbModel.AddForeignKey("source_id", "feedback_sources(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding feedback-source foreign key: %w", err)
	}
	if err = fbModel.AddForeignKey("type_id", "feedback_types(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding feedback-type foreign key: %w", err)
	}
	if err = fbModel.AddForeignKey("station_id", "stations(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding feedback-station foreign key: %w", err)
	}
	if err = fbModel.AddForeignKey("line_id", "lines(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding feedback-line foreign key: %w", err)
	}
	if err = fbModel.AddForeignKey("direction_id", "directions(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding feedback-direction foreign key: %w", err)
	}

	if err = db.AutoMigrate(&Train{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `Train`: %w", err)
	}

	if err = db.AutoMigrate(&ScheduleEventSource{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `ScheduleEventSource`: %w", err)
	}
	if err = db.AutoMigrate(&ScheduleEvent{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `ScheduleEvent`: %w", err)
	}

	sEvent := db.Model(&ScheduleEvent{})
	if err = sEvent.AddForeignKey("event_type_id", "schedule_event_sources(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding schedule-event/event_type foreign key: %w", err)
	}
	if err = sEvent.AddForeignKey("destination_id", "stations(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding schedule-event/destination foreign key: %w", err)
	}
	if err = sEvent.AddForeignKey("next_station_id", "stations(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding schedule-event/next_station foreign key: %w", err)
	}

	if err = db.AutoMigrate(&RealTimeEventDetail{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `RealTimeEventDetail`: %w", err)
	}

	rtedTable := db.Model(&RealTimeEventDetail{})
	if err = rtedTable.AddForeignKey("schedule_event_id", "schedule_events(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding real-time-event-detail/schedule_event: %w", err)
	}
	if err = rtedTable.AddForeignKey("train_id", "trains(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding real-time-event-detail/train: %w", err)
	}

	if err = db.AutoMigrate(&StaticEventDetail{}).Error; err != nil {
		return fmt.Errorf("failed auto-migrating table `StaticEventDetail`: %w", err)
	}
	if err = db.Model(&StaticEventDetail{}).AddForeignKey("schedule_event_id", "schedule_events(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return fmt.Errorf("failed adding static-event-detail/schedule-event foreign key: %w", err)
	}

	return nil
}
