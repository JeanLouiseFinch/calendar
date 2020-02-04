package entities

import "time"

//Event - nashe vse
type Event struct {
	Start       time.Time
	End         time.Time
	Owner       string
	Title       string
	Description string
}
