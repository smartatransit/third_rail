package seed

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/third_rail/pkg/models"
)

func Seed(db *gorm.DB) {

	db.LogMode(false)

	//Event sources
	UpsertEventSource(db, "MARTA_StaticSchedule", "MARTA's Trip Schedule")
	UpsertEventSource(db, "MARTA_RealTime", "MARTA's Real Time API")

	//Directions
	log.Info("Creating Directions")

	northbound := UpsertDirection(db, "Northbound", []string{"NB", "N", "North"})
	southbound := UpsertDirection(db, "Southbound", []string{"SB", "S", "South"})
	eastbound := UpsertDirection(db, "Eastbound", []string{"EB", "E", "East"})
	westbound := UpsertDirection(db, "Westbound", []string{"WB", "W", "West"})

	//Lines
	log.Info("Creating Lines")

	gold := UpsertLine(db, "Gold", []string{"Gold"}, []models.Direction{northbound, southbound})
	red := UpsertLine(db, "Red", nil, []models.Direction{northbound, southbound})
	blue := UpsertLine(db, "Blue", nil, []models.Direction{eastbound, westbound})
	green := UpsertLine(db, "Green", nil, []models.Direction{eastbound, westbound})

	//Stations
	log.Info("Creating Stations")

	//Gold
	_ = UpsertStation(db, "Doraville", "dnh0f5v6mxzj", "Doraville Station", nil, []models.Line{gold})
	_ = UpsertStation(db, "Chamblee", "dnh0c94gqrm2", "Chamblee Station", nil, []models.Line{gold})
	_ = UpsertStation(db, "Brookhaven", "dnh08u1fpng1", "Brookhaven Station", []string{"Brookhaven-Oglethorpe Station"}, []models.Line{gold})
	_ = UpsertStation(db, "Lenox", "dnh0837g7frm", "Lenox Station", nil, []models.Line{gold})

	//Red
	_ = UpsertStation(db, "North Springs", "dnh107scwx23", "North Springs Station", nil, []models.Line{red})
	_ = UpsertStation(db, "Sandy Springs", "dnh10939s87f", "Sandy Springs Station", nil, []models.Line{red})
	_ = UpsertStation(db, "Dunwoody", "dnh0bxnr3hcj", "Dunwoody Station", nil, []models.Line{red})
	_ = UpsertStation(db, "Medical Center", "dnh0bt32f0zr", "Medical Center Station", nil, []models.Line{red})
	_ = UpsertStation(db, "Buckhead", "dnh084sbc4fj", "Buckhead Station", nil, []models.Line{red})

	//Blue
	_ = UpsertStation(db, "Indian Creek", "dnh0579u6fcg", "Indian Creek Station", nil, []models.Line{blue})
	_ = UpsertStation(db, "Kensington", "dnh04u1s7ycg", "Kensington Station", nil, []models.Line{blue})
	_ = UpsertStation(db, "Avondale", "dnh04heebfzp", "Avondale Station", nil, []models.Line{blue})
	_ = UpsertStation(db, "Decatur", "dnh01u9cru4h", "Decatur Station", nil, []models.Line{blue})
	_ = UpsertStation(db, "East Lake", "dnh016v1ynv9", "East Lake Station", nil, []models.Line{blue})
	_ = UpsertStation(db, "West Lake", "dn5bn2sgc1bc", "West Lake Station", nil, []models.Line{blue})
	_ = UpsertStation(db, "H. E. Holmes", "dn5bjbfgcmr3", "Hamilton E. Holmes Station", []string{"Hamilton E. Holmes", "Hamilton E Holmes", "Hamilton E Holmes Station"}, []models.Line{blue})

	//Green
	_ = UpsertStation(db, "Bankhead", "dn5bnu0epkt1", "Bankhead Station", nil, []models.Line{green})

	//Multi-line
	_ = UpsertStation(db, "Lindbergh Center", "dn5bnu0epkt1", "Lindbergh Center Station", []string{"Lindbergh", "Lindbergh Station"}, []models.Line{gold, red})
	_ = UpsertStation(db, "Arts Center", "dn5bnu0epkt1", "Arts Center Station", nil, []models.Line{gold, red})
	_ = UpsertStation(db, "Midtown", "dn5bptxy8r41", "Midtown Station", nil, []models.Line{gold, red})
	_ = UpsertStation(db, "North Avenue", "dn5bpsp70th7", "North Avenue Station", []string{"North Ave", "North Ave Station"}, []models.Line{gold, red})
	_ = UpsertStation(db, "Civic Center", "dn5bpep496h7", "Civic Center Station", nil, []models.Line{gold, red})
	_ = UpsertStation(db, "Peachtree Center", "dn5bp9qxs9nh", "Peachtree Center Station", nil, []models.Line{gold, red})
	_ = UpsertStation(db, "Five Points", "dn5bp8ezwy4k", "Five Points Station", []string{"5 points"}, []models.Line{gold, red, blue, green})
	_ = UpsertStation(db, "Garnett", "djgzzxbb2xkd", "Garnett Station", nil, []models.Line{gold, red})
	_ = UpsertStation(db, "West End", "djgzzjeb581t", "West End Station", nil, []models.Line{gold, red})
	_ = UpsertStation(db, "Oakland City", "djgzyf5dfsmn", "Oakland City Station", []string{"Oakland"}, []models.Line{gold, red})
	_ = UpsertStation(db, "Lakewood", "djgzwz0c6suf", "Lakewood/Fort McPherson Station", []string{"Lakewood", "Lakewood Station", "Ft. Mcpherson"}, []models.Line{gold, red})
	_ = UpsertStation(db, "East Point", "djgzwdb63g2k", "East Point Station", nil, []models.Line{gold, red})
	_ = UpsertStation(db, "College Park", "djgzqq4k3j73", "College Park Station", nil, []models.Line{gold, red})
	_ = UpsertStation(db, "Airport", "djgzqkhjse84", "Airport Station", nil, []models.Line{gold, red})
	_ = UpsertStation(db, "Ashby", "dn5bp11qp0s9", "Ashby Station", nil, []models.Line{blue, green})
	_ = UpsertStation(db, "Vine City", "dn5bp34zmh5s", "Vine City Station", nil, []models.Line{blue, green})
	_ = UpsertStation(db, "Omni Dome", "dn5bp90pezjh", "Omni/Dome/GWCC/State Farm/CNN Center Station", []string{"Omni", "Omni Dome", "Omni Dome Station", "Georgia Dome", "CNN", "State Farm Arena", "Phillips Arena", "Georgia World Congress"}, []models.Line{blue, green})
	_ = UpsertStation(db, "Georgia State", "dn5bp8pgdtcf", "Georgia State Station", []string{"GSU"}, []models.Line{blue, green})
	_ = UpsertStation(db, "King Memorial", "dn5bpbp8fe6s", "King Memorial Station", []string{"MLK"}, []models.Line{blue, green})
	_ = UpsertStation(db, "Inman Park", "dnh0092w0nxh", "Inman Park-Reynoldstown Station", []string{"Inman Park Station", "Reynoldstown"}, []models.Line{blue, green})
	_ = UpsertStation(db, "Edgewood-Candler Park", "dnh00f1wzrc6", "Edgewood-Candler Park Station", []string{"Edgewood", "Candler", "Edgewood Candler Park", "Edgewood Candler Park Station"}, []models.Line{blue, green})

}

func UpsertEventSource(db *gorm.DB, name, description string) models.ScheduleEventSource {
	var eventSource models.ScheduleEventSource

	db.FirstOrCreate(&eventSource, &models.ScheduleEventSource{
		Name:        name,
		Description: description,
	})

	return eventSource
}

func UpsertDirection(db *gorm.DB, name string, aliases []string) models.Direction {
	direction := models.Direction{Name: name}

	if aliases != nil {
		var namedAliases []models.Alias

		for _, a := range aliases {
			namedAliases = append(namedAliases, models.Alias{Alias: a})
		}

		direction.Aliases = namedAliases
	}

	db.Save(&direction)

	return direction
}

func UpsertLine(db *gorm.DB, name string, aliases []string, directions []models.Direction) models.Line {
	line := models.Line{Name: name}

	if aliases != nil {
		var namedAliases []models.Alias

		for _, a := range aliases {
			namedAliases = append(namedAliases, models.Alias{Alias: a})
		}

		line.Aliases = namedAliases
	}

	db.Save(&line)
	db.Model(&line).Association("Directions").Append(directions)

	return line
}

func UpsertStation(db *gorm.DB, name, location, description string, aliases []string, lines []models.Line) models.Station {
	var station models.Station

	if aliases != nil {
		var namedAliases []models.Alias

		for _, a := range aliases {
			namedAliases = append(namedAliases, models.Alias{Alias: a})
		}

		station = models.Station{
			Name:    name,
			Aliases: namedAliases,
		}

	} else {
		station = models.Station{
			Name: name,
		}
	}

	db.Save(&station)

	db.Model(&station).Association("Lines").Append(lines)

	stationDetail := models.StationDetail{
		StationID:   station.ID,
		Description: description,
		Location:    location,
	}

	db.Save(&stationDetail)

	return station
}
