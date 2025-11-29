package events

import (
	"context"

	eventsDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
	tasksService "github.com/aevula/interview-hustlers-calendar/internal/services/tasks"
)

type EventsService interface {
	Create(ctx context.Context, cmd CreateEventCommand) (eventsDomain.Event, error)
	Update(ctx context.Context, cmd UpdateEventCommand) (eventsDomain.Event, error)
	Delete(ctx context.Context, cmd DeleteEventCommand) error
	List(ctx context.Context, cmd ListEventsCommand) ([]eventsDomain.Event, error)
}

type eventsService struct {
	repo         EventRepository
	tasksService tasksService.TasksService
}

type EventRepository interface {
	Create(ctx context.Context, event eventsDomain.Event) (eventsDomain.Event, error)
	Update(ctx context.Context, event eventsDomain.Event) (eventsDomain.Event, error)
	Delete(ctx context.Context, id eventsDomain.EventID) error
	FindAllByStartAt(ctx context.Context, filter eventsDomain.StartAtFilter) ([]eventsDomain.Event, error)
}

func NewEventsService(repo EventRepository, tasksService tasksService.TasksService) EventsService {
	return &eventsService{repo: repo, tasksService: tasksService}
}
