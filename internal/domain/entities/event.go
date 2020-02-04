package entities

import "time"

//Event - nashe vse
type Event struct {
	start       time.Time
	end         time.Time
	owner       string
	title       string
	description string
}
