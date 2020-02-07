package interfaces

import (
	"context"

	"github.com/JeanLouiseFinch/calendar/internal/domain/entities"
)

type Storage interface {
	AddEvent(ctx context.Context, e *entities.Event) error
	DeleteEvent(ctx context.Context, id uint) error
	GetEventByID(ctx context.Context, id uint) (*entities.Event, error)
	GetEventsByTitle(ctx context.Context, title string) ([]*entities.Event, error)
	GetEventsByOwner(ctx context.Context, owner string) ([]*entities.Event, error)
	EditEvent(ctx context.Context, id uint, e *entities.Event) error
	GetEventsByTimeRange(ctx context.Context, timeRange string) error
	String() string
}
