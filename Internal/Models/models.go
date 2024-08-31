package models

import "time"

type Event struct {
	EventId      int
	Title        string
	CreationDate time.Time
	AlertType    int
	Priority     int
	DueDate      time.Time
	DueTime      time.Time
	Note         string
}
