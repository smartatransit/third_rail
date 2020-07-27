package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/gomarta"
	"time"
)

const FETCH_SCHEDULE_EVENTS_BY_STATION_REALTIME string = `
SELECT DISTINCT ON(se.next_station_id, rted.train_id) se.id, rted.id
FROM schedule_events se
LEFT JOIN real_time_event_details rted on se.id = rted.schedule_event_id
WHERE event_type_id = 2 AND next_station_id = ?
ORDER BY se.next_station_id, rted.train_id, rted.created_at DESC
`
const FETCH_SCHEDULE_EVENTS_BY_STATION_STATIC string = ""

type ScheduleEvent struct {
	ID            uint `json:"-" gorm:"primary_key"`
	EventTypeID   uint `json:"-"`
	EventType     ScheduleEventSource
	DestinationID uint `json:"-"`
	Destination   Station
	NextArrival   time.Time
	NextStationID uint `json:"-"`
	NextStation   Station
	DirectionID   uint `json:"-"`
	Direction     Direction
	ExpiresAt     time.Time

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

func (se ScheduleEvent) LoadFromGoMartaEvent(gmTrain gomarta.Train, db *gorm.DB) {

}

func GetScheduleEventsByStationRealTime(stationId int, db *gorm.DB) (scheduleEvents []ScheduleEvent, realTimeDetails []RealTimeEventDetail, err error) {
	rows, err := db.Raw(FETCH_SCHEDULE_EVENTS_BY_STATION_REALTIME, stationId).Rows()

	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var scheduleId string
		var realTimeDetailId string
		var scheduleEvent ScheduleEvent
		var realTimeDetail RealTimeEventDetail

		rows.Scan(&scheduleId, &realTimeDetailId)

		db.Preload("EventType").Preload("Destination").Preload("NextStation").Preload("Direction").First(&scheduleEvent, scheduleId)
		db.Preload("Train").First(&realTimeDetail, realTimeDetailId)

		scheduleEvents = append(scheduleEvents, scheduleEvent)
		realTimeDetails = append(realTimeDetails, realTimeDetail)
	}

	return
}
