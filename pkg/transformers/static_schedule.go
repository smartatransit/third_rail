package transformers

import (
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/third_rail/pkg/schemas/marta_schemas"
)

type StaticScheduleTransformer struct {
	Schedules struct{
		Weekday []marta_schemas.StationSchedule
		Saturday []marta_schemas.StationSchedule
		Sunday []marta_schemas.StationSchedule
	}
}

func NewStaticScheduleTransformer() StaticScheduleTransformer {
	_, err := parseCsv("data/location/stations.csv")

	if err != nil {
		log.Fatal("Unable to load static station data: %s", err)
	}

	return StaticScheduleTransformer{}
}

func (lt StaticScheduleTransformer) GetSchedule(schedule, stationName string) (sortedStationLocations []marta_schemas.StationLocation) {
	return
}

func parseScheduleData(stationData [][]string) (stationLocations []marta_schemas.StationLocation) {
	for i, _ := range stationData[0] {
		station := marta_schemas.StationLocation{
			StationName: stationData[0][i],
			Location:    stationData[1][i],
			Distance:    0,
		}
		stationLocations = append(stationLocations, station)
	}

	return
}
