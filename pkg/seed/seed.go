package seed

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/third_rail/pkg/models"
)

func Seed(db *gorm.DB) {

	db.LogMode(true)

	//Directions
	log.Info("Creating Directions")

	northbound := UpsertDirection(db, "Northbound")
	southbound := UpsertDirection(db, "Southbound")
	eastbound := UpsertDirection(db, "Eastbound")
	westbound := UpsertDirection(db, "Westbound")

	//Lines
	log.Info("Creating Lines")

	gold := UpsertLine(db, "Gold", []models.Direction{northbound, southbound})
	red := UpsertLine(db, "Red", []models.Direction{northbound, southbound})
	blue := UpsertLine(db, "Blue", []models.Direction{eastbound, westbound})
	green := UpsertLine(db, "Green", []models.Direction{eastbound, westbound})

	//Stations
	log.Info("Creating Stations")

	//Gold
	_ = UpsertStation(db, "Doraville", "dnh0f5v6mxzj", "Doraville Station", []models.Line{gold})
	_ = UpsertStation(db, "Chamblee", "dnh0c94gqrm2", "Chamblee Station", []models.Line{gold})
	_ = UpsertStation(db, "Brookhaven", "dnh08u1fpng1", "Brookhaven Station", []models.Line{gold})
	_ = UpsertStation(db, "Lenox", "dnh0837g7frm", "Lenox Station", []models.Line{gold})

	//Red
	_ = UpsertStation(db, "North Springs", "dnh107scwx23", "North Springs Station", []models.Line{red})
	_ = UpsertStation(db, "Sandy Springs", "dnh10939s87f", "Sandy Springs Station", []models.Line{red})
	_ = UpsertStation(db, "Dunwoody", "dnh0bxnr3hcj", "Dunwoody Station", []models.Line{red})
	_ = UpsertStation(db, "Medical Center", "dnh0bt32f0zr", "Medical Center Station", []models.Line{red})
	_ = UpsertStation(db, "Buckhead", "dnh084sbc4fj", "Buckhead Station", []models.Line{red})

	//Blue
	_ = UpsertStation(db, "Indian Creek", "dnh0579u6fcg", "Indian Creek Station", []models.Line{blue})
	_ = UpsertStation(db, "Kensington", "dnh04u1s7ycg", "Kensington Station", []models.Line{blue})
	_ = UpsertStation(db, "Avondale", "dnh04heebfzp", "Avondale Station", []models.Line{blue})
	_ = UpsertStation(db, "Decatur", "dnh01u9cru4h", "Decatur Station", []models.Line{blue})
	_ = UpsertStation(db, "East Lake", "dnh016v1ynv9", "East Lake Station", []models.Line{blue})
	_ = UpsertStation(db, "West Lake", "dn5bn2sgc1bc", "West Lake Station", []models.Line{blue})
	_ = UpsertStation(db, "H. E. Holmes", "dn5bjbfgcmr3", "Hamilton E. Holmes Station", []models.Line{blue})

	//Green
	_ = UpsertStation(db, "Bankhead", "dn5bnu0epkt1", "Bankhead Station", []models.Line{green})

	//Multi-line
	_ = UpsertStation(db, "Lindebergh Center", "dn5bnu0epkt1", "Lindbergh Center Station", []models.Line{gold, red})
	_ = UpsertStation(db, "Arts Center", "dn5bnu0epkt1", "Arts Center Station", []models.Line{gold, red})
	_ = UpsertStation(db, "Midtown", "dn5bnu0epkt1", "Midtown Station", []models.Line{gold, red})
	_ = UpsertStation(db, "North Avenue", "dn5bnu0epkt1", "North Avenue Station", []models.Line{gold, red})
	_ = UpsertStation(db, "Civic Center", "dn5bnu0epkt1", "Civic Center Station", []models.Line{gold, red})
	_ = UpsertStation(db, "Peachtree Center", "dn5bnu0epkt1", "Peachtree Center Station", []models.Line{gold, red})
	_ = UpsertStation(db, "Five Points", "dn5bp8ezwy4k", "Five Points Stations", []models.Line{gold, red, blue, green})
	_ = UpsertStation(db, "Garnett", "djgzzxbb2xkd", "Garnett Station", []models.Line{gold, red})
	_ = UpsertStation(db, "West End", "djgzzjeb581t", "West End Station", []models.Line{gold, red})
	_ = UpsertStation(db, "Oakland City", "djgzyf5dfsmn", "Oakland City Station", []models.Line{gold, red})
	_ = UpsertStation(db, "Lakewood", "djgzwz0c6suf", "Lakewood/Fort McPherson Station", []models.Line{gold, red})
	_ = UpsertStation(db, "East Point", "djgzwdb63g2k", "East Point Station", []models.Line{gold, red})
	_ = UpsertStation(db, "College Park", "djgzqq4k3j73", "College Park Station", []models.Line{gold, red})
	_ = UpsertStation(db, "Airport", "djgzqkhjse84", "Airport Station", []models.Line{gold, red})
	_ = UpsertStation(db, "Ashby", "dn5bp11qp0s9", "Ashby Station", []models.Line{blue, green})
	_ = UpsertStation(db, "Vine City", "dn5bp34zmh5s", "Vine City Station", []models.Line{blue, green})
	_ = UpsertStation(db, "Omni Dome", "dn5bp90pezjh", "Omni/Dome/GWCC/State Farm/CNN Center Station", []models.Line{blue, green})
	_ = UpsertStation(db, "Georgia State", "dn5bp8pgdtcf", "Georgia State Station", []models.Line{blue, green})
	_ = UpsertStation(db, "King Memorial", "dn5bpbp8fe6s", "King Memorial Station", []models.Line{blue, green})
	_ = UpsertStation(db, "Inman Park", "dnh0092w0nxh", "Inman Station", []models.Line{blue, green})
	_ = UpsertStation(db, "Edgewood-Candler Park", "dnh00f1wzrc6", "Edgewood-Candler Park Station", []models.Line{blue, green})

}

func UpsertDirection(db *gorm.DB, name string) models.Direction {
	var direction models.Direction
	db.FirstOrCreate(&direction, &models.Direction{
		Name: name,
	})

	return direction
}

func UpsertLine(db *gorm.DB, name string, directions []models.Direction) models.Line {
	var line models.Line
	db.FirstOrCreate(&line, &models.Line{
		Name: name,
	})
	db.Model(&line).Association("Directions").Append(directions)

	return line
}

func UpsertStation(db *gorm.DB, name, location, description string, lines []models.Line) models.Station {
	var station models.Station
	db.FirstOrCreate(&station, &models.Station{
		Name: name,
	})
	db.Model(&station).Association("Lines").Append(lines)

	var stationDetail models.StationDetail
	db.FirstOrCreate(&stationDetail, &models.StationDetail{
		StationID:   station.ID,
		Description: description,
		Location:    location,
	})

	return station
}
