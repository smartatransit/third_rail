package transformers

import (
	"encoding/csv"
	"github.com/mmcloughlin/geohash"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/third_rail/pkg/schemas/marta_schemas"
	"math"
	"os"
	"sort"
)

type LocationTransformer struct {
	Stations []marta_schemas.StationLocation
}

func NewLocationTransformer() LocationTransformer {
	stations, err := parseCsv("data/location/stations.csv")

	if err != nil {
		log.Fatal("Unable to load static station data: %s", err)
	}

	return LocationTransformer{parseStationData(stations)}
}

func (lt LocationTransformer) GetNearestLocations(latitude, longitude float64) (sortedStationLocations []marta_schemas.StationLocation) {
	for _, locationStation := range lt.Stations {
		stationLat, stationLong := geohash.Decode(locationStation.Location)
		locationStation.Distance = calculateDistance(latitude, longitude, stationLat, stationLong)
		sortedStationLocations = append(sortedStationLocations, locationStation)
	}

	sort.Slice(sortedStationLocations[:], func(i, j int) bool {
		return sortedStationLocations[i].Distance < sortedStationLocations[j].Distance
	})

	return
}

func parseStationData(stationData [][]string) (stationLocations []marta_schemas.StationLocation) {
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


func parseCsv(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Unable to read input file "+fileName, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+fileName, err)
	}

	return records, err
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func calculateDistance(startingLatitude, startingLongitude, endingLatitude, endingLongitude float64) float64 {
	var latitude1, longitude1, latitude2, longitude2, radius, meterFeet float64

	latitude1 = startingLatitude * math.Pi / 180
	longitude1 = startingLongitude * math.Pi / 180
	latitude2 = endingLatitude * math.Pi / 180
	longitude2 = endingLongitude * math.Pi / 180

	radius = 6378100
	meterFeet = 3.2808399

	h := hsin(latitude2-latitude1) + math.Cos(latitude1)*math.Cos(latitude2)*hsin(longitude2-longitude1)

	return 2 * radius * math.Asin(math.Sqrt(h)) * meterFeet
}
