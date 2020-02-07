package nosql

import (
	"context"
	"testing"
	"time"

	"github.com/JeanLouiseFinch/calendar/internal/domain/entities"
	"github.com/JeanLouiseFinch/calendar/internal/domain/errors"
)

func TestAddEvent(t *testing.T) {
	events := []*entities.Event{
		&entities.Event{Title: "Event1", Owner: "me", Description: "wow1", Start: time.Now(), End: time.Now().AddDate(0, 0, 1)},
		&entities.Event{Title: "Event2", Owner: "me", Description: "wow2", Start: time.Now().AddDate(0, 0, 2), End: time.Now().AddDate(0, 0, 3)},
		&entities.Event{Title: "Event3", Owner: "me", Description: "wow3", Start: time.Now().AddDate(0, 0, 2), End: time.Now().AddDate(0, 0, 3)},
	}
	ss := NewStorage()
	_, err := ss.AddEvent(context.Background(), events[0])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = ss.AddEvent(context.Background(), events[1])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = ss.AddEvent(context.Background(), events[2])
	if err != errors.ErrEventBusy {
		t.Error("Expected ErrEventBusy, got ", err)
	}
}

func TestGetEventsByTimeRange(t *testing.T) {
	events := []*entities.Event{
		&entities.Event{Title: "Event1", Owner: "me", Description: "wow1", Start: time.Now(), End: time.Now().AddDate(0, 0, 1)},
		&entities.Event{Title: "Event2", Owner: "me", Description: "wow2", Start: time.Now().AddDate(0, 0, 2), End: time.Now().AddDate(0, 0, 3)},
		&entities.Event{Title: "Event3", Owner: "me", Description: "wow3", Start: time.Now().AddDate(0, 0, 3), End: time.Now().AddDate(0, 0, 9)},
		&entities.Event{Title: "Event3", Owner: "me", Description: "wow3", Start: time.Now().AddDate(0, 0, 10), End: time.Now().AddDate(0, 0, 30)},
		&entities.Event{Title: "Event3", Owner: "me", Description: "wow3", Start: time.Now().AddDate(0, 3, 10), End: time.Now().AddDate(0, 5, 30)},
	}
	ss := NewStorage()
	for _, val := range events {
		_, err := ss.AddEvent(context.Background(), val)
		if err != nil {
			t.Error("Expected nil, got ", err)
		}
	}
	dd, err := ss.GetEventsByTimeRange(context.Background(), entities.TimeRangeDay)
	if err != nil && len(dd) != 1 {
		t.Error("Expected nil, got ", err)
	}
	ww, err := ss.GetEventsByTimeRange(context.Background(), entities.TimeRangeWeek)
	if err != nil && len(ww) != 3 {
		t.Error("Expected nil, got ", err)
	}
	mm, err := ss.GetEventsByTimeRange(context.Background(), entities.TimeRangeMounth)
	if err != nil && len(mm) != 4 {
		t.Error("Expected nil, got ", err)
	}
	yy, err := ss.GetEventsByTimeRange(context.Background(), entities.TimeRangeYear)
	if err != nil && len(yy) != 5 {
		t.Error("Expected nil, got ", err)
	}
	_, err = ss.GetEventsByTimeRange(context.Background(), "dasfjdsfgsdf")
	if err != errors.ErrTimeRange {
		t.Error("Expected ErrTimeRange, got ", err)
	}
}
func TestDeleteEvent(t *testing.T) {
	events := []*entities.Event{
		&entities.Event{Title: "Event1", Owner: "me", Description: "wow1", Start: time.Now(), End: time.Now().AddDate(0, 0, 1)},
		&entities.Event{Title: "Event2", Owner: "me", Description: "wow2", Start: time.Now().AddDate(0, 0, 2), End: time.Now().AddDate(0, 0, 3)},
	}
	storage := NewStorage()
	id, err := storage.AddEvent(context.Background(), events[0])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = storage.AddEvent(context.Background(), events[1])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	err = storage.DeleteEvent(context.Background(), id)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	err = storage.DeleteEvent(context.Background(), id)
	if err != errors.ErrEventNotFound {
		t.Error("Expected ErrEventNotFound, got ", err)
	}
}
func TestGetEventById(t *testing.T) {
	events := []*entities.Event{
		&entities.Event{Title: "Event1", Owner: "me", Description: "wow1", Start: time.Now(), End: time.Now().AddDate(0, 0, 1)},
		&entities.Event{Title: "Event2", Owner: "me", Description: "wow2", Start: time.Now().AddDate(0, 0, 2), End: time.Now().AddDate(0, 0, 3)},
	}
	storage := NewStorage()
	id, err := storage.AddEvent(context.Background(), events[0])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = storage.AddEvent(context.Background(), events[1])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = storage.GetEventByID(context.Background(), id)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = storage.GetEventByID(context.Background(), id+999)
	if err != errors.ErrEventNotFound {
		t.Error("Expected ErrEventNotFound, got ", err)
	}
}

func TestGetEventByTitle(t *testing.T) {
	events := []*entities.Event{
		&entities.Event{Title: "Event1", Owner: "me", Description: "wow1", Start: time.Now(), End: time.Now().AddDate(0, 0, 1)},
		&entities.Event{Title: "Event2", Owner: "me", Description: "wow2", Start: time.Now().AddDate(0, 0, 2), End: time.Now().AddDate(0, 0, 3)},
	}
	storage := NewStorage()
	_, err := storage.AddEvent(context.Background(), events[0])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = storage.AddEvent(context.Background(), events[1])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	es, err := storage.GetEventsByTitle(context.Background(), "Event2")
	if err != nil && es[0].Title != "Event2" {
		t.Error("Expected nil, got ", err)
	}
	_, err = storage.GetEventsByTitle(context.Background(), "Event3")
	if err != errors.ErrEventNotFound {
		t.Error("Expected ErrEventNotFound, got ", err)
	}
}

func TestGetEventByOwner(t *testing.T) {
	events := []*entities.Event{
		&entities.Event{Title: "Event1", Owner: "me", Description: "wow1", Start: time.Now(), End: time.Now().AddDate(0, 0, 1)},
		&entities.Event{Title: "Event2", Owner: "me", Description: "wow2", Start: time.Now().AddDate(0, 0, 2), End: time.Now().AddDate(0, 0, 3)},
	}
	storage := NewStorage()
	_, err := storage.AddEvent(context.Background(), events[0])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = storage.AddEvent(context.Background(), events[1])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	es, err := storage.GetEventsByOwner(context.Background(), "me")
	if err != nil && es[0].Title != "me" && es[1].Title != "me" {
		t.Error("Expected nil, got ", err)
	}
	_, err = storage.GetEventsByOwner(context.Background(), "not me")
	if err != errors.ErrEventNotFound {
		t.Error("Expected ErrEventNotFound, got ", err)
	}
}

func TestEditEvent(t *testing.T) {
	events := []*entities.Event{
		&entities.Event{Title: "Event1", Owner: "me", Description: "wow1", Start: time.Now(), End: time.Now().AddDate(0, 0, 1)},
		&entities.Event{Title: "Event2", Owner: "me", Description: "wow2", Start: time.Now().AddDate(0, 0, 2), End: time.Now().AddDate(0, 0, 3)},
		&entities.Event{Title: "Event3", Owner: "bu", Description: "wow3", Start: time.Now().AddDate(0, 1, 2), End: time.Now().AddDate(0, 1, 3)},
	}
	storage := NewStorage()
	id, err := storage.AddEvent(context.Background(), events[0])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = storage.AddEvent(context.Background(), events[1])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	err = storage.EditEvent(context.Background(), id, events[2])
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	ev, err := storage.GetEventByID(context.Background(), id)
	if err != nil && ev.Title != "Event3" {
		t.Error("Expected nil, got ", err)
	}

	err = storage.EditEvent(context.Background(), id+999, events[2])
	if err != errors.ErrEventNotFound {
		t.Error("Expected ErrEventNotFound, got ", err)
	}
}
