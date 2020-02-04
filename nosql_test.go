package main

import (
	"context"
	"testing"
	"time"

	"github.com/JeanLouiseFinch/calendar/internal/adapters/nosql"
	"github.com/JeanLouiseFinch/calendar/internal/domain/entities"
	"github.com/JeanLouiseFinch/calendar/internal/domain/errors"
)

func TestAddEvent(t *testing.T) {
	events := []*entities.Event{
		&entities.Event{Title: "Event1", Owner: "me", Description: "wow1", Start: time.Now(), End: time.Now().AddDate(0, 0, 1)},
		&entities.Event{Title: "Event2", Owner: "me", Description: "wow2", Start: time.Now().AddDate(0, 0, 2), End: time.Now().AddDate(0, 0, 3)},
		&entities.Event{Title: "Event3", Owner: "me", Description: "wow3", Start: time.Now().AddDate(0, 0, 2), End: time.Now().AddDate(0, 0, 3)},
	}
	storage := nosql.NewStorage()
	_, err := storage.AddEvent(context.Background(), events[0])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = storage.AddEvent(context.Background(), events[1])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = storage.AddEvent(context.Background(), events[2])
	if err != errors.ErrEventBusy {
		t.Error("Expected ErrEventBusy, got ", err)
	}
}
