package events

import (
	"context"
)

type DeleteEventCommand struct {
	ID int
}

func (s *eventsService) Delete(ctx context.Context, cmd DeleteEventCommand) error {
	return s.repo.Delete(ctx, cmd.ID)
}
