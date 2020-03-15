package transformers

import (
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/gomarta"
	"github.com/smartatransit/third_rail/pkg/schemas/marta_schemas"
	"github.com/smartatransit/third_rail/pkg/validators"
	"strings"
)

type EventTransformer struct {
	MEV validators.MartaEntitiesValidator
}

func NewEventTransformer(mev validators.MartaEntitiesValidator) (et EventTransformer) {
	et.MEV = mev

	return
}

func (et EventTransformer) GetStation(event gomarta.Train) marta_schemas.Station {
	direction, _ := et.MEV.Coerce(validators.MARTA_DIRECTIONS, string(event.Direction))
	line, _ := et.MEV.Coerce(validators.MARTA_LINES, string(event.Line))
	station, err := et.MEV.Coerce(validators.MARTA_STATIONS, event.Station)

	if err != nil {
		log.Printf("Coercion miss: %s", err)
	}

	return marta_schemas.Station{
		Direction: direction,
		Line:      line,
		Name:      station,
	}
}

func (et EventTransformer) GetSchedule(event gomarta.Train) marta_schemas.Schedule {
	destination, destErr := et.MEV.Coerce(validators.MARTA_STATIONS, event.Destination)

	if destErr != nil {
		log.Printf("Coercion miss: %s", destErr)
	}

	station, statErr := et.MEV.Coerce(validators.MARTA_STATIONS, event.Station)

	if statErr != nil {
		log.Printf("Coercion miss: %s", statErr)
	}

	return marta_schemas.Schedule{
		Destination:    destination,
		EventTime:      event.EventTime,
		NextArrival:    event.NextArrival,
		NextStation:    station,
		TrainID:        event.TrainID,
		WaitingSeconds: event.WaitingSeconds,
		WaitingTime:    event.WaitingTime,
	}
}

func FilterByLine(events []gomarta.Train, lineFilter string) (lineEvents []gomarta.Train) {
	for _, event := range events {
		eventLine := strings.ToUpper(string(event.Line))
		if eventLine == strings.ToUpper(lineFilter) {
			lineEvents = append(lineEvents, event)
		}
	}

	return
}

func FilterByStation(events []gomarta.Train, stationFilter string) (stationEvents []gomarta.Train) {
	for _, event := range events {
		eventStation := strings.ToUpper(event.Station)
		if eventStation == strings.ToUpper(stationFilter) {
			stationEvents = append(stationEvents, event)
		}
	}

	return
}
