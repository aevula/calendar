package events

import (
	"context"
	"time"

	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
)

type UpdateEventCommand struct {
	ID          int
	Title       string
	Description string
	StartAt     time.Time
	Duration    time.Duration
	NotifyAt    time.Time
}

func (s *eventsService) Update(ctx context.Context, cmd UpdateEventCommand) (domain.Event, error) {
	event := domain.Event{
		ID:          cmd.ID,
		Title:       cmd.Title,
		Description: cmd.Description,
		StartAt:     cmd.StartAt,
		Duration:    cmd.Duration,
		NotifyAt:    cmd.NotifyAt,
	}

	event.EnsureNotifyAt()

	return s.repo.Update(ctx, event)
}
