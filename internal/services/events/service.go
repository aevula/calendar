package events

import (
	"context"

	"github.com/aevula/interview-hustlers-calendar/internal/repository"
)

type EventsService interface {
	Create(ctx context.Context, cmd CreateEventCommand) (repository.Event, error)
	Update(ctx context.Context, cmd UpdateEventCommand) (repository.Event, error)
	Delete(ctx context.Context, cmd DeleteEventCommand) error
	All(ctx context.Context) ([]repository.Event, error)
}

type eventsService struct {
	repo repository.EventRepository
}

func NewEventsService(repo repository.EventRepository) EventsService {
	return &eventsService{repo: repo}
}
