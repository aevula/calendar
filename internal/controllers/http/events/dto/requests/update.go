package requests

import (
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/services/events"
)

type UpdateEventRequest struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	Duration    string    `json:"duration"`
	NotifyAt    time.Time `json:"notify_at"`
}

func (r UpdateEventRequest) ToCommand() (cmd events.UpdateEventCommand, err error) {
	duration, err := time.ParseDuration(r.Duration)
	if err != nil {
		return
	}

	cmd = events.UpdateEventCommand{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		StartAt:     r.StartAt,
		Duration:    duration,
		NotifyAt:    r.NotifyAt,
	}
	return
}
