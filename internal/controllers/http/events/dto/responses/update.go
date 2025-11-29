package responses

import (
	"time"

	eventsDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
)

type UpdateEventResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	StartAt     time.Time `json:"start_at"`
	Duration    string    `json:"duration"`
	NotifyAt    time.Time `json:"notify_at"`
}

func ToUpdateEventResponse(event eventsDomain.Event) UpdateEventResponse {
	return UpdateEventResponse{
		ID:          int(event.ID),
		Title:       event.Title,
		Description: event.Description,
		UserId:      int(event.UserId),
		StartAt:     event.StartAt,
		Duration:    event.Duration.String(),
		NotifyAt:    event.NotifyAt,
	}
}
