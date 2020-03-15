package marta_schemas

type Schedule struct {
	Destination    string `json:"destination"`
	EventTime      string `json:"event_time"`
	NextArrival    string `json:"next_arrival"`
	NextStation    string `json:"next_station"`
	TrainID        string `json:"train_id"`
	WaitingSeconds string `json:"waiting_seconds"`
	WaitingTime    string `json:"waiting_time"`
}

type Schedules []Schedule
