package requests

import (
	"github.com/aevula/interview-hustlers-calendar/internal/services/events"
)

type DeleteEventRequest struct {
	ID int `json:"id"`
}

func (r DeleteEventRequest) ToCommand() events.DeleteEventCommand {
	return events.DeleteEventCommand{
		ID: r.ID,
	}
}
