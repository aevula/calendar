package responses

import (
	"time"

	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
)

type CreateEventResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	StartAt     time.Time `json:"start_at"`
	Duration    string    `json:"duration"`
	NotifyAt    time.Time `json:"notify_at"`
}

func ToCreateEventResponse(event domain.Event) CreateEventResponse {
	return CreateEventResponse{
		ID:          int(event.ID),
		Title:       event.Title,
		Description: event.Description,
		UserId:      int(event.UserId),
		StartAt:     event.StartAt,
		Duration:    event.Duration.String(),
		NotifyAt:    event.NotifyAt,
	}
}
