package requests

import (
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/services/events"
)

type ListEventsRequest struct {
	StartAt StartAtRange `json:"start_at"`
}

type StartAtRange struct {
	GTE time.Time `json:"gte"`
	LTE time.Time `json:"lte"`
}

func (r ListEventsRequest) ToCommand() events.ListEventsCommand {
	return events.ListEventsCommand{
		StartAt: events.StartAtRange{
			GTE: r.StartAt.GTE,
			LTE: r.StartAt.LTE,
		},
	}
}
