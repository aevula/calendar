package events

import (
	"context"

	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
)

type EventsService interface {
	Create(ctx context.Context, cmd CreateEventCommand) (domain.Event, error)
	Update(ctx context.Context, cmd UpdateEventCommand) (domain.Event, error)
	Delete(ctx context.Context, cmd DeleteEventCommand) error
	All(ctx context.Context) ([]domain.Event, error)
}

type eventsService struct {
	repo EventRepository
}

type EventRepository interface {
	Create(ctx context.Context, event domain.Event) (domain.Event, error)
	Update(ctx context.Context, event domain.Event) (domain.Event, error)
	Delete(ctx context.Context, id domain.EventID) error
	All(ctx context.Context) ([]domain.Event, error)
}

func NewEventsService(repo EventRepository) EventsService {
	return &eventsService{repo: repo}
}
