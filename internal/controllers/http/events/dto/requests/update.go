package requests

import (
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/services/events"
)

type UpdateEventRequest struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	StartAt     time.Time     `json:"start_at"`
	Duration    time.Duration `json:"duration"`
	NotifyAt    time.Time     `json:"notify_at"`
}

func (r UpdateEventRequest) ToCommand() events.UpdateEventCommand {
	return events.UpdateEventCommand{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		StartAt:     r.StartAt,
		Duration:    r.Duration,
		NotifyAt:    r.NotifyAt,
	}
}
