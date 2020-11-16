package seed

import (
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/third_rail/pkg/models"
)

func Seed(db *gorm.DB, log *log.Logger) error {
	var err error

	//Event sources
	if _, err = InsertEventSource(db, "MARTA_StaticSchedule", "MARTA's Trip Schedule"); err != nil {
		return err
	}
	if _, err = InsertEventSource(db, "MARTA_RealTime", "MARTA's Real Time API"); err != nil {
		return err
	}

	//Directions
	log.Info("Creating Directions")

	northbound, err := InsertDirection(db, "Northbound", []string{"NB", "N", "North"})
	if err != nil {
		return err
	}
	southbound, err := InsertDirection(db, "Southbound", []string{"SB", "S", "South"})
	if err != nil {
		return err
	}
	eastbound, err := InsertDirection(db, "Eastbound", []string{"EB", "E", "East"})
	if err != nil {
		return err
	}
	westbound, err := InsertDirection(db, "Westbound", []string{"WB", "W", "West"})
	if err != nil {
		return err
	}

	//Lines
	log.Info("Creating Lines")

	gold, err := InsertLine(db, "Gold", []string{"Gold"}, []models.Direction{northbound, southbound})
	if err != nil {
		return err
	}
	red, err := InsertLine(db, "Red", nil, []models.Direction{northbound, southbound})
	if err != nil {
		return err
	}
	blue, err := InsertLine(db, "Blue", nil, []models.Direction{eastbound, westbound})
	if err != nil {
		return err
	}
	green, err := InsertLine(db, "Green", nil, []models.Direction{eastbound, westbound})
	if err != nil {
		return err
	}

	//Stations
	log.Info("Creating Stations")

	//Gold
	_, err = InsertStation(db, "Doraville", "dnh0f5v6mxzj", "Doraville Station", nil, []models.Line{gold})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Chamblee", "dnh0c94gqrm2", "Chamblee Station", nil, []models.Line{gold})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Brookhaven", "dnh08u1fpng1", "Brookhaven Station", []string{"Brookhaven-Oglethorpe Station"}, []models.Line{gold})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Lenox", "dnh0837g7frm", "Lenox Station", nil, []models.Line{gold})
	if err != nil {
		return err
	}

	//Red
	_, err = InsertStation(db, "North Springs", "dnh107scwx23", "North Springs Station", nil, []models.Line{red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Sandy Springs", "dnh10939s87f", "Sandy Springs Station", nil, []models.Line{red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Dunwoody", "dnh0bxnr3hcj", "Dunwoody Station", nil, []models.Line{red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Medical Center", "dnh0bt32f0zr", "Medical Center Station", nil, []models.Line{red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Buckhead", "dnh084sbc4fj", "Buckhead Station", nil, []models.Line{red})
	if err != nil {
		return err
	}

	//Blue
	_, err = InsertStation(db, "Indian Creek", "dnh0579u6fcg", "Indian Creek Station", nil, []models.Line{blue})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Kensington", "dnh04u1s7ycg", "Kensington Station", nil, []models.Line{blue})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Avondale", "dnh04heebfzp", "Avondale Station", nil, []models.Line{blue})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Decatur", "dnh01u9cru4h", "Decatur Station", nil, []models.Line{blue})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "East Lake", "dnh016v1ynv9", "East Lake Station", nil, []models.Line{blue})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "West Lake", "dn5bn2sgc1bc", "West Lake Station", nil, []models.Line{blue})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "H. E. Holmes", "dn5bjbfgcmr3", "Hamilton E. Holmes Station", []string{"Hamilton E. Holmes", "Hamilton E Holmes", "Hamilton E Holmes Station"}, []models.Line{blue})
	if err != nil {
		return err
	}

	//Green
	_, err = InsertStation(db, "Bankhead", "dn5bnu0ejdq9", "Bankhead Station", nil, []models.Line{green})
	if err != nil {
		return err
	}

	//Multi-line
	_, err = InsertStation(db, "Lindbergh Center", "dnh02jhnbebq", "Lindbergh Center Station", []string{"Lindbergh", "Lindbergh Station"}, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Arts Center", "dn5bpxphcqh3", "Arts Center Station", nil, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Midtown", "dn5bptxy8r41", "Midtown Station", nil, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "North Avenue", "dn5bpsp70th7", "North Avenue Station", []string{"North Ave", "North Ave Station"}, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Civic Center", "dn5bpep496h7", "Civic Center Station", nil, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Peachtree Center", "dn5bp9qxs9nh", "Peachtree Center Station", nil, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Five Points", "dn5bp8ezwy4k", "Five Points Station", []string{"5 points"}, []models.Line{gold, red, blue, green})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Garnett", "djgzzxbb2xkd", "Garnett Station", nil, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "West End", "djgzzjeb581t", "West End Station", nil, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Oakland City", "djgzyf5dfsmn", "Oakland City Station", []string{"Oakland"}, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Lakewood", "djgzwz0c6suf", "Lakewood/Fort McPherson Station", []string{"Lakewood", "Lakewood Station", "Ft. Mcpherson"}, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "East Point", "djgzwdb63g2k", "East Point Station", nil, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "College Park", "djgzqq4k3j73", "College Park Station", nil, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Airport", "djgzqkhjse84", "Airport Station", nil, []models.Line{gold, red})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Ashby", "dn5bp11qp0s9", "Ashby Station", nil, []models.Line{blue, green})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Vine City", "dn5bp34zmh5s", "Vine City Station", nil, []models.Line{blue, green})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Omni Dome", "dn5bp90pezjh", "Omni/Dome/GWCC/State Farm/CNN Center Station", []string{"Omni", "Omni Dome", "Omni Dome Station", "Georgia Dome", "CNN", "State Farm Arena", "Phillips Arena", "Georgia World Congress"}, []models.Line{blue, green})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Georgia State", "dn5bp8pgdtcf", "Georgia State Station", []string{"GSU"}, []models.Line{blue, green})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "King Memorial", "dn5bpbp8fe6s", "King Memorial Station", []string{"MLK"}, []models.Line{blue, green})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Inman Park", "dnh0092w0nxh", "Inman Park-Reynoldstown Station", []string{"Inman Park Station", "Reynoldstown"}, []models.Line{blue, green})
	if err != nil {
		return err
	}
	_, err = InsertStation(db, "Edgewood-Candler Park", "dnh00f1wzrc6", "Edgewood-Candler Park Station", []string{"Edgewood", "Candler", "Edgewood Candler Park", "Edgewood Candler Park Station"}, []models.Line{blue, green})
	if err != nil {
		return err
	}

	return nil
}

func InsertEventSource(db *gorm.DB, name, description string) (models.ScheduleEventSource, error) {
	var err error
	var eventSource models.ScheduleEventSource

	err = db.FirstOrCreate(&eventSource, &models.ScheduleEventSource{
		Name:        name,
		Description: description,
	}).Error
	if err != nil {
		return eventSource, fmt.Errorf("failed creating event source `%s`: %w", name, err)
	}

	return eventSource, nil
}

func InsertDirection(db *gorm.DB, name string, aliases []string) (models.Direction, error) {
	var err error
	direction := models.Direction{Name: name}

	if aliases != nil {
		var namedAliases []models.Alias

		for _, a := range aliases {
			namedAliases = append(namedAliases, models.Alias{Alias: a})
		}

		direction.Aliases = namedAliases
	}

	err = db.Save(&direction).Error
	if err != nil {
		return direction, fmt.Errorf("failed creating direction `%s`: %w", name, err)
	}

	return direction, nil
}

func InsertLine(db *gorm.DB, name string, aliases []string, directions []models.Direction) (models.Line, error) {
	var err error
	line := models.Line{Name: name}

	if aliases != nil {
		var namedAliases []models.Alias

		for _, a := range aliases {
			namedAliases = append(namedAliases, models.Alias{Alias: a})
		}

		line.Aliases = namedAliases
	}

	err = db.Save(&line).Error
	if err != nil {
		return line, err
	}
	err = db.Model(&line).Association("Directions").Append(directions).Error
	if err != nil {
		return line, err
	}

	return line, nil
}

func InsertStation(db *gorm.DB, name, location, description string, aliases []string, lines []models.Line) (models.Station, error) {
	var err error
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

	err = db.Save(&station).Error
	if err != nil {
		return station, fmt.Errorf("failed creating station `%s`: %w", name, err)
	}

	err = db.Model(&station).Association("Lines").Append(lines).Error
	if err != nil {
		return station, fmt.Errorf("failed creating station `%s`: %w", name, err)
	}

	stationDetail := models.StationDetail{
		StationID:   station.ID,
		Description: description,
		Location:    location,
	}

	err = db.Save(&stationDetail).Error
	if err != nil {
		return station, fmt.Errorf("failed creating station `%s`: %w", name, err)
	}

	return station, nil
}
