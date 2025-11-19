package events

import (
	"context"
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/repository"
)

type UpdateEventCommand struct {
	ID           int
	Title        string
	Description  string
	StartAt      time.Time
	Duration     time.Duration
	NotifyOffset time.Duration
}

func (s *eventsService) Update(ctx context.Context, cmd UpdateEventCommand) (repository.Event, error) {
	event := repository.Event{
		ID:           cmd.ID,
		Title:        cmd.Title,
		Description:  cmd.Description,
		StartAt:      cmd.StartAt,
		Duration:     cmd.Duration,
		NotifyOffset: cmd.NotifyOffset,
	}
	return s.repo.Update(ctx, event)
}
