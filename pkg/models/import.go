package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/gomarta"
	"strconv"
	"time"
)

func ImportLiveEvents(db *gorm.DB, events []gomarta.Train) {

	db = db.Set("gorm:auto_preload", true)

	var realTimeEventSource ScheduleEventSource
	db.Where(&ScheduleEventSource{Name: "MARTA_RealTime"}).First(&realTimeEventSource)

	for _, event := range events {
		now := time.Now()
		nextArrival, _ := time.Parse("03:04:05 PM", event.NextArrival)
		adjustedNextArrival := time.Date(now.Year(), now.Month(), now.Day(), nextArrival.Hour(), nextArrival.Minute(), nextArrival.Second(), nextArrival.Nanosecond(), nextArrival.Location())

		destination, destErr := FindStationByName(event.Destination, db)

		if destErr != nil {
			log.Error(destErr)
		}

		nextStation, nextErr := FindStationByName(event.Station, db)

		if nextErr != nil {
			log.Error(fmt.Sprintf("Unable to find %s", event.Station))
			log.Error(nextErr)
		}

		direction, dirErr := FindDirectionByName(string(event.Direction), db)

		if dirErr != nil {
			log.Error(dirErr)
		}

		expiresAt := time.Now().Add(time.Minute * 20)

		scheduleEvent := ScheduleEvent{
			EventTypeID:   realTimeEventSource.ID,
			DestinationID: destination.ID,
			NextArrival:   adjustedNextArrival,
			NextStationID: nextStation.ID,
			DirectionID:   direction.ID,
			ExpiresAt:     expiresAt,
		}

		db.Create(&scheduleEvent)

		trainID, _ := strconv.ParseUint(event.TrainID, 10, 32)
		train := Train{
			ID: uint(trainID),
		}

		db.FirstOrCreate(&train)

		eventTime, _ := time.Parse("1/2/2006 3:04:05 PM", event.EventTime)
		waitSeconds, _ := strconv.Atoi(event.WaitingSeconds)

		realTimeEvent := RealTimeEventDetail{
			ScheduleEventID: scheduleEvent.ID,
			EventTime:       eventTime,
			TrainID:         uint(trainID),
			WaitingSeconds:  waitSeconds,
			WaitingTime:     event.WaitingTime,
		}

		db.Create(&realTimeEvent)
	}

}
