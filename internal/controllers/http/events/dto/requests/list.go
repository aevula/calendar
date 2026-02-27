package requests

import (
	"time"

	eventsService "github.com/aevula/interview-hustlers-calendar/internal/services/events"
)

type ListEventsRequest struct {
	StartAt StartAtRange `json:"start_at"`
}

type StartAtRange struct {
	GTE time.Time `json:"gte"`
	LTE time.Time `json:"lte"`
}

func (r ListEventsRequest) ToCommand() eventsService.ListEventsCommand {
	return eventsService.ListEventsCommand{
		StartAt: eventsService.StartAtRange{
			GTE: r.StartAt.GTE,
			LTE: r.StartAt.LTE,
		},
	}
}
