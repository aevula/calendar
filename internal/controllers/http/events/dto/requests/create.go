package requests

import (
	"time"

	eventsService "github.com/aevula/interview-hustlers-calendar/internal/services/events"
)

type CreateEventRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	StartAt     time.Time `json:"start_at"`
	Duration    string    `json:"duration"`
	NotifyAt    time.Time `json:"notify_at"`
}

func (r CreateEventRequest) ToCommand() (cmd eventsService.CreateEventCommand, err error) {
	duration, err := time.ParseDuration(r.Duration)
	if err != nil {
		return
	}

	cmd = eventsService.CreateEventCommand{
		Title:       r.Title,
		Description: r.Description,
		UserId:      r.UserId,
		StartAt:     r.StartAt,
		Duration:    duration,
		NotifyAt:    r.NotifyAt,
	}
	return
}
