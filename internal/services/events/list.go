package events

import (
	"context"
	"time"

	eventsDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
)

type ListEventsCommand struct {
	StartAt StartAtRange
}

type StartAtRange struct {
	GTE time.Time
	LTE time.Time
}

func (s *eventsService) List(ctx context.Context, cmd ListEventsCommand) ([]eventsDomain.Event, error) {
	filter := eventsDomain.StartAtFilter{
		GTE: cmd.StartAt.GTE,
		LTE: cmd.StartAt.LTE,
	}
	return s.repo.FindAllByStartAt(ctx, filter)
}
