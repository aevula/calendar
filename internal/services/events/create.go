package events

import (
	"context"
	"time"

	eventsDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
	tasksService "github.com/aevula/interview-hustlers-calendar/internal/services/tasks"
)

type CreateEventCommand struct {
	Title       string
	Description string
	UserId      int
	StartAt     time.Time
	Duration    time.Duration
	NotifyAt    time.Time
}

func (s *eventsService) Create(ctx context.Context, cmd CreateEventCommand) (eventsDomain.Event, error) {
	event := eventsDomain.Event{
		Title:       cmd.Title,
		Description: cmd.Description,
		UserId:      eventsDomain.UserID(cmd.UserId),
		StartAt:     cmd.StartAt,
		Duration:    cmd.Duration,
		NotifyAt:    cmd.NotifyAt,
	}

	event.EnsureNotifyAt()

	event, err := s.repo.Create(ctx, event)
	if err != nil {
		return event, err
	}

	taskCmd := tasksService.CreateTaskCommand{
		Name:    "send_events_notifications",
		Queue:   "send_events_notifications_queue",
		Payload: map[string]any{"eventID": event.ID},
	}

	_, err = s.tasksService.Create(ctx, taskCmd)
	return event, err
}
