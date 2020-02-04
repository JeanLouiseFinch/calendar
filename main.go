package main

import (
	"context"
	"time"

	"github.com/JeanLouiseFinch/calendar/internal/adapters/nosql"
	"github.com/JeanLouiseFinch/calendar/internal/domain/entities"
	"github.com/JeanLouiseFinch/calendar/internal/domain/error"
)

func main() {
	events := []*entities.Event {
		&entities.Event{Title: "Event1",Owner: "me",Description: "wow1",Start: time.Now(),End: time.Now().AddDate(0,0,1)},
		&entities.Event{Title: "Event2",Owner: "me",Description: "wow2",Start: time.Now().AddDate(0,0,2),End: time.Now().AddDate(0,0,3)},
		&entities.Event{Title: "Event3",Owner: "not me",Description: "wow3",time.Now().AddDate(0,0,2),End: time.Now().AddDate(0,0,3)},
		&entities.Event{Title: "Event4",Owner: "me",Description: "wow4",time.Now().AddDate(0,0,2),End: time.Now().AddDate(0,0,3)},
		&entities.Event{Title: "Event5",Owner: "not me",Description: "wow5",time.Now().AddDate(0,0,2),End: time.Now().AddDate(0,0,3)}
	}
	storage := nosql.NewStorage()
	err := storage.AddEvent(context.Background(),events[0])
	if err != nil {
		panic(err)
	}
}
