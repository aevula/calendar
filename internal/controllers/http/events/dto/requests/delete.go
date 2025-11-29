package requests

import (
	eventsService "github.com/aevula/interview-hustlers-calendar/internal/services/events"
)

type DeleteEventRequest struct {
	ID int `json:"id"`
}

func (r DeleteEventRequest) ToCommand() eventsService.DeleteEventCommand {
	return eventsService.DeleteEventCommand{
		ID: r.ID,
	}
}
