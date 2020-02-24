package transformers

import (
	"github.com/smartatransit/third_rail/pkg/schemas/marta_schemas"
	"log"
)

type StaticScheduleTransformer struct {
	Stations []marta_schemas.StationLocation
}

func NewStaticScheduleTransformer() StaticScheduleTransformer {
	stations, err := parseCsv("data/location/stations.csv")

	if err != nil {
		log.Panic("OH SHIT OH SHIT OH SHIT NO CSV FOUND OH SHIT")
	}

	return StaticScheduleTransformer{parseStationData(stations)}
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
