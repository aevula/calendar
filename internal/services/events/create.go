package events

import (
	"context"
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/repository"
)

type CreateEventCommand struct {
	Title       string
	Description string
	UserId      int
	StartAt     time.Time
	Duration    time.Duration
	NotifyAt    time.Time
}

func (s *eventsService) Create(ctx context.Context, cmd CreateEventCommand) (repository.Event, error) {
	event := repository.Event{
		Title:       cmd.Title,
		Description: cmd.Description,
		UserId:      cmd.UserId,
		StartAt:     cmd.StartAt,
		Duration:    cmd.Duration,
		NotifyAt:    cmd.NotifyAt,
	}
	return s.repo.Create(ctx, event)
}
