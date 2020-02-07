package nosql

import (
	"context"
	"fmt"
	"sync"
	"time"

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
	return &Storage{events: make(map[uint]*entities.Event, 0), mutex: &sync.Mutex{}}
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
func (s *Storage) GetEventsByTimeRange(ctx context.Context, timeRange string) ([]*entities.Event, error) {
	var duration time.Time

	switch timeRange {
	case entities.TimeRangeDay:
		duration = time.Now().AddDate(0, 0, 1)
	case entities.TimeRangeWeek:
		duration = time.Now().AddDate(0, 0, 7)
	case entities.TimeRangeMounth:
		duration = time.Now().AddDate(0, 1, 0)
	case entities.TimeRangeYear:
		duration = time.Now().AddDate(1, 0, 0)
	default:
		return nil, errors.ErrTimeRange
	}
	now := time.Now()
	result := make([]*entities.Event, 0, 0)
	for _, value := range s.events {
		if (value.Start.After(now) && value.Start.Before(duration)) || (value.Start == duration || value.Start == now) || (value.End.After(now) && value.End.Before(duration)) || (value.End == now || value.End == duration) {
			result = append(result, value)
		}
	}
	if len(result) > 0 {
		return result, nil
	}
	return nil, errors.ErrEventNotFound
}

// String - realizacia interface
func (s *Storage) String() string {
	result := "Storage:\n-------\n"
	for key, value := range s.events {
		result += fmt.Sprintf("\tEvent:%s[%d] by %s\n\t\t%s. Time: %v-%v\n---\n", value.Title, key, value.Owner, value.Description, value.Start, value.End)
	}
	return result
}
