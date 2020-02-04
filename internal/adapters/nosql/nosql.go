package nosql

import (
	"context"
	"fmt"
	"sync"

	"github.com/JeanLouiseFinch/calendar/internal/domain/entities"
	"github.com/JeanLouiseFinch/calendar/internal/domain/errors"
)

// Storage - nasha hranilka
type Storage struct {
	events    map[uint]*entities.Event
	currIndex uint
	mutex     *sync.Mutex
}

// NewStorage - sozdaem hranilishe
func NewStorage() *Storage {
	return &Storage{
		events: make(map[uint]*entities.Event, 0),
		mutex:  &sync.Mutex{},
	}
}

// AddEvent - realizacia interface
func (s *Storage) AddEvent(ctx context.Context, e *entities.Event) (uint, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	was := false
	for _, value := range s.events {
		if (e.Start.After(value.Start) && e.Start.Before(value.End)) || (e.End.After(value.Start) && e.End.Before(value.End) || e.Start.Equal(value.Start) || e.Start.Equal(value.End) || e.End.Equal(value.Start) || e.End.Equal(value.End)) {
			was = true
		}
	}
	if !was {
		s.currIndex++
		s.events[s.currIndex] = e
		return s.currIndex, nil
	}
	return 0, errors.ErrEventBusy
}

// DeleteEvent - realizacia interface
func (s *Storage) DeleteEvent(ctx context.Context, id uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.events[id]; ok {
		delete(s.events, id)
		return nil
	}
	return errors.ErrEventNotFound
}

// GetEventByID - realizacia interface
func (s *Storage) GetEventByID(ctx context.Context, id uint) (*entities.Event, error) {
	if _, ok := s.events[id]; ok {
		return s.events[id], nil
	}
	return nil, errors.ErrEventNotFound
}

// GetEventsByTitle - realizacia interface
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

// GetEventsByOwner - realizacia interface
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

// EditEvent - realizacia interface
func (s *Storage) EditEvent(ctx context.Context, id uint, e *entities.Event) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.events[id]; ok {
		s.events[id] = e
		return nil
	}
	return errors.ErrEventNotFound
}

// String - realizacia interface
func (s *Storage) String() string {
	result := "Storage:\n-------\n"
	for key, value := range s.events {
		result += fmt.Sprintf("\tEvent:%s[%d] by %s\n\t\t%s. Time: %v-%v\n---\n", value.Title, key, value.Owner, value.Description, value.Start, value.End)
	}
	return result
}
