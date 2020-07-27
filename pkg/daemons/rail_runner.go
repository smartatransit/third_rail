package daemons

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/models"
	"strconv"
	"time"
)

type RailRunner struct {
	MartaClient clients.MartaClient
	DB          *gorm.DB
}

func (rr RailRunner) Start() {
	log.Info("Starting Rail Runner")

	ticker := time.NewTicker(15 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				log.Info("Rail Runner finished.")
				return
			case t := <-ticker.C:
				log.Infof("Fetch MARTA events at %v", t)
				rr.ImportEvents()
				log.Info("Fetch finished")
			}
		}
	}()
}

func (rr RailRunner) ImportEvents() {
	events, err := rr.MartaClient.GetTrains()
	rr.DB = rr.DB.Set("gorm:auto_preload", true)
	if err != nil {
		log.Error(err)
	}

	var realTimeEventSource models.ScheduleEventSource
	rr.DB.Where(&models.ScheduleEventSource{Name: "MARTA_RealTime"}).First(&realTimeEventSource)

	for _, event := range events {
		now := time.Now()
		nextArrival, _ := time.Parse("03:04:05 PM", event.NextArrival)
		adjustedNextArrival := time.Date(now.Year(), now.Month(), now.Day(), nextArrival.Hour(), nextArrival.Minute(), nextArrival.Second(), nextArrival.Nanosecond(), nextArrival.Location())

		destination, destErr := models.FindStationByName(event.Destination, rr.DB)

		if destErr != nil {
			log.Error(destErr)
		}

		nextStation, nextErr := models.FindStationByName(event.Station, rr.DB)

		if nextErr != nil {
			log.Error(fmt.Sprintf("Unable to find %s", event.Station))
			log.Error(nextErr)
		}

		direction, dirErr := models.FindDirectionByName(string(event.Direction), rr.DB)

		if dirErr != nil {
			log.Error(dirErr)
		}

		expiresAt := time.Now().Add(time.Minute * 20)

		scheduleEvent := models.ScheduleEvent{
			EventTypeID:   realTimeEventSource.ID,
			DestinationID: destination.ID,
			NextArrival:   adjustedNextArrival,
			NextStationID: nextStation.ID,
			DirectionID:   direction.ID,
			ExpiresAt:     expiresAt,
		}

		rr.DB.Create(&scheduleEvent)

		trainID, _ := strconv.ParseUint(event.TrainID, 10, 32)
		train := models.Train{
			ID: uint(trainID),
		}

		rr.DB.FirstOrCreate(&train)

		eventTime, _ := time.Parse("1/2/2006 3:04:05 PM", event.EventTime)
		waitSeconds, _ := strconv.Atoi(event.WaitingSeconds)

		realTimeEvent := models.RealTimeEventDetail{
			ScheduleEventID: scheduleEvent.ID,
			EventTime:       eventTime,
			TrainID:         uint(trainID),
			WaitingSeconds:  waitSeconds,
			WaitingTime:     event.WaitingTime,
		}

		rr.DB.Create(&realTimeEvent)
	}

}
