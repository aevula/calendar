package requests

import (
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/services/events"
)

type UpdateEventRequest struct {
	ID           int
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	StartAt      time.Time     `json:"start_at"`
	Duration     time.Duration `json:"duration"`
	NotifyOffset time.Duration `json:"notify_offset"`
}

func (r UpdateEventRequest) ToCommand() events.UpdateEventCommand {
	return events.UpdateEventCommand{
		Title:        r.Title,
		Description:  r.Description,
		StartAt:      r.StartAt,
		Duration:     r.Duration,
		NotifyOffset: r.NotifyOffset,
	}
}
