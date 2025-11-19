package events

import (
	"context"

	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
)

func (s *eventsService) All(ctx context.Context) ([]domain.Event, error) {
	return s.repo.All(ctx)
}
