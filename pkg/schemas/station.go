package schemas

type Station struct {
	Direction string `json:"direction"`
	Line      string `json:"line"`
	Name      string `json:"name"`
}

type StationLocation struct {
	StationName string  `json:"station_name"`
	Location    string  `json:"location"`
	Distance    float64 `json:"distance"`
}
