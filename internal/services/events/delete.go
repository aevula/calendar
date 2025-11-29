package events

import (
	"context"

	eventsDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
)

type DeleteEventCommand struct {
	ID int
}

func (s *eventsService) Delete(ctx context.Context, cmd DeleteEventCommand) error {
	return s.repo.Delete(ctx, eventsDomain.EventID(cmd.ID))
}
