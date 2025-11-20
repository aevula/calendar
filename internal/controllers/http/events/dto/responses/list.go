package responses

import (
	"time"

	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
)

type listEventResponse struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UserId      int        `json:"user_id"`
	StartAt     time.Time  `json:"start_at"`
	Duration    string     `json:"duration"`
	NotifyAt    time.Time  `json:"notify_at"`
	NotifiedAt  *time.Time `json:"notified_at"`
}

type ListEventsResponse = []listEventResponse

func ToListEventsResponse(events []domain.Event) ListEventsResponse {
	list := make([]listEventResponse, 0, len(events))
	for i := range events {
		var notified_at *time.Time
		if !events[i].NotifiedAt.IsZero() {
			notified_at = &events[i].NotifiedAt
		}

		list = append(list, listEventResponse{
			ID:          int(events[i].ID),
			Title:       events[i].Title,
			Description: events[i].Description,
			UserId:      int(events[i].UserId),
			StartAt:     events[i].StartAt,
			Duration:    events[i].Duration.String(),
			NotifyAt:    events[i].NotifyAt,
			NotifiedAt:  notified_at,
		})
	}

	return list
}
