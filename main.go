package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JeanLouiseFinch/calendar/internal/adapters/nosql"
	"github.com/JeanLouiseFinch/calendar/internal/domain/entities"
)

func main() {
	events := []*entities.Event{
		&entities.Event{Title: "Event1", Owner: "me", Description: "wow1", Start: time.Now(), End: time.Now().AddDate(0, 0, 1)},
		&entities.Event{Title: "Event2", Owner: "me", Description: "wow2", Start: time.Now().AddDate(0, 0, 2), End: time.Now().AddDate(0, 0, 3)},
		&entities.Event{Title: "Event3", Owner: "not me", Description: "wow3", Start: time.Now().AddDate(0, 0, 2), End: time.Now().AddDate(0, 0, 3)},
		&entities.Event{Title: "Event4", Owner: "me", Description: "wow4", Start: time.Now().AddDate(0, 1, 2), End: time.Now().AddDate(0, 1, 3)},
		&entities.Event{Title: "Event5", Owner: "not me", Description: "wow5", Start: time.Now().AddDate(1, 0, 2), End: time.Now().AddDate(1, 0, 3)},
	}

	storage := nosql.NewStorage()
	id, err := storage.AddEvent(context.Background(), events[0])
	if err != nil {
		fmt.Println(err)
	}
	id2, err := storage.AddEvent(context.Background(), events[1])
	if err != nil {
		fmt.Println(err)
	}
	_, err = storage.AddEvent(context.Background(), events[2])
	if err != nil {
		fmt.Println(err)
	}
	_, err = storage.AddEvent(context.Background(), events[3])
	if err != nil {
		fmt.Println(err)
	}
	_, err = storage.AddEvent(context.Background(), events[4])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(storage)

	fmt.Printf("DELETE\n")
	err = storage.DeleteEvent(context.Background(), id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(storage)
	fmt.Printf("EDIT\n")
	e := &entities.Event{Title: "EventX", Owner: "me", Description: "wow1", Start: time.Now().AddDate(2, 0, 1), End: time.Now().AddDate(2, 2, 1)}
	err = storage.EditEvent(context.Background(), id2, e)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(storage)
}
