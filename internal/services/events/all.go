package events

import (
	"context"

	"github.com/aevula/interview-hustlers-calendar/internal/repository"
)

func (s *eventsService) All(ctx context.Context) ([]repository.Event, error) {
	return s.repo.All(ctx)
}
