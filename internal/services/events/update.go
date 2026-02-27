package events

import (
	"context"
	"time"

	eventsDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
)

type UpdateEventCommand struct {
	ID          int
	Title       string
	Description string
	StartAt     time.Time
	Duration    time.Duration
	NotifyAt    time.Time
}

func (s *eventsService) Update(ctx context.Context, cmd UpdateEventCommand) (eventsDomain.Event, error) {
	event := eventsDomain.Event{
		ID:          eventsDomain.EventID(cmd.ID),
		Title:       cmd.Title,
		Description: cmd.Description,
		StartAt:     cmd.StartAt,
		Duration:    cmd.Duration,
		NotifyAt:    cmd.NotifyAt,
	}

	event.EnsureNotifyAt()

	return s.repo.Update(ctx, event)
}
