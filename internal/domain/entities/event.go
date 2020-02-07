package entities

import "time"

const (
	TimeRangeYear   = "year"
	TimeRangeMounth = "mounth"
	TimeRangeWeek   = "week"
	TimeRangeDay    = "day"
)

//Event - nashe vse
type Event struct {
	Start       time.Time
	End         time.Time
	Owner       string
	Title       string
	Description string
}
