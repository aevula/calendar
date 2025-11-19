package requests

import (
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/services/events"
)

type CreateEventRequest struct {
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	UserId       int           `json:"user_id"`
	StartAt      time.Time     `json:"start_at"`
	Duration     time.Duration `json:"duration"`
	NotifyOffset time.Duration `json:"notify_offset"`
}

func (r CreateEventRequest) ToCommand() events.CreateEventCommand {
	return events.CreateEventCommand{
		Title:        r.Title,
		Description:  r.Description,
		UserId:       r.UserId,
		StartAt:      r.StartAt,
		Duration:     r.Duration,
		NotifyOffset: r.NotifyOffset,
	}
}
