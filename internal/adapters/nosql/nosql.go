package nosql

import (
	"context"
	"fmt"
	"sync"

	"github.com/JeanLouiseFinch/calendar/internal/domain/entities"
	"github.com/JeanLouiseFinch/calendar/internal/domain/errors"
)

type Storage struct {
	events    map[int]*entities.Event
	currIndex int
	mutex     *sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		events: make(map[int]*entities.Event, 0),
		mutex:  &sync.Mutex{},
	}
}

func (s *Storage) AddEvent(ctx context.Context, e *entities.Event) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	was := false
	for _, value := range s.events {
		if (e.Start.After(value.Start) && e.Start.Before(value.End)) || (e.End.After(value.Start) && e.End.Before(value.End)) {
			was = true
		}
	}
	if !was {
		s.currIndex++
		s.events[s.currIndex] = e
		return nil
	}
	return errors.ErrEventBusy
}
func (s *Storage) DeleteEvent(ctx context.Context, id uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.events[id]; ok {
		delete(s.events[id])
		return nil
	}
	return errors.ErrEventNotFound
}
func (s *Storage) GetEventByID(ctx context.Context, id uint) (*entities.Event, error) {
	if _, ok := s.events[id]; ok {
		return s.events[id]
	}
	return errors.ErrEventNotFound
}
func (s *Storage) GetEventsByTitle(ctx context.Context, title string) ([]*entities.Event, error) {
	result := make([]*entities.Event, 0, 0)
	for _, value := range s.events {
		if value.Title == title {
			result = append(result, value)
		}
	}
	if len(result) > 0 {
		return result, nil
	}
	return nil, errors.ErrEventNotFound
}

func (s *Storage) GetEventsByOwner(ctx context.Context, owner string) ([]*entities.Event, error) {
	result := make([]*entities.Event, 0, 0)
	for _, value := range s.events {
		if value.Owner == owner {
			result = append(result, value)
		}
	}
	if len(result) > 0 {
		return result, nil
	}
	return nil, errors.ErrEventNotFound
}
func (s *Storage) EditEvent(ctx context.Context, id uint, e *entities.Event) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.events[id]; ok {
		s.events[id] = e
		return nil
	}
	return errors.ErrEventNotFound
}
func (s *Storage) String() string {
	result := "Storage:\n-------\n"
	for _, value := range s.events {
		result += fmt.Printf("\tEvent:%s by %s\n\t\t%s. Time: %v-%v\n---\n", value.Title, value.Owner, value.Description, value.Start, value.End)
	}
}
