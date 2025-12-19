package events

import (
	"context"
	"time"

	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
)

type CreateEventCommand struct {
	Title       string
	Description string
	UserId      int
	StartAt     time.Time
	Duration    time.Duration
	NotifyAt    time.Time
}

func (s *eventsService) Create(ctx context.Context, cmd CreateEventCommand) (domain.Event, error) {
	event := domain.Event{
		Title:       cmd.Title,
		Description: cmd.Description,
		UserId:      domain.UserID(cmd.UserId),
		StartAt:     cmd.StartAt,
		Duration:    cmd.Duration,
		NotifyAt:    cmd.NotifyAt,
	}

	event.EnsureNotifyAt()

	return s.repo.Create(ctx, event)
}
