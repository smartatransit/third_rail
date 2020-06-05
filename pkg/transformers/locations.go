package transformers

import (
	"github.com/mmcloughlin/geohash"
	"github.com/smartatransit/third_rail/pkg/models"
	"math"
	"sort"
)

func SortStationsByDistance(latitude, longitude float64, stations []models.Station) (sortedStations []models.Station){
	for _, locationStation := range stations {
		stationLat, stationLong := geohash.Decode(locationStation.Detail.Location)
		locationStation.Detail.Distance = calculateDistance(latitude, longitude, stationLat, stationLong)
		sortedStations = append(sortedStations, locationStation)
	}

	sort.Slice(sortedStations[:], func(i, j int) bool {
		return sortedStations[i].Detail.Distance < sortedStations[j].Detail.Distance
	})

	return
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
