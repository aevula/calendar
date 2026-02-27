package events

import (
	"time"

	eventsDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
)

type Event struct {
	ID          int
	Title       string
	Description string
	UserId      int
	StartAt     time.Time
	Duration    time.Duration
	NotifyAt    time.Time
	NotifiedAt  *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func FromDomain(event eventsDomain.Event) Event {
	var notified_at *time.Time
	if !event.NotifiedAt.IsZero() {
		notified_at = &event.NotifiedAt
	}

	return Event{
		ID:          int(event.ID),
		Title:       event.Title,
		Description: event.Description,
		UserId:      int(event.UserId),
		StartAt:     event.StartAt,
		Duration:    event.Duration,
		NotifyAt:    event.NotifyAt,
		NotifiedAt:  notified_at,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}
}

func (event Event) ToDomain() eventsDomain.Event {
	var notifiedAt time.Time
	if event.NotifiedAt != nil {
		notifiedAt = *event.NotifiedAt
	}

	return eventsDomain.Event{
		ID:          eventsDomain.EventID(event.ID),
		Title:       event.Title,
		Description: event.Description,
		UserId:      eventsDomain.UserID(event.UserId),
		StartAt:     event.StartAt,
		Duration:    event.Duration,
		NotifyAt:    event.NotifyAt,
		NotifiedAt:  notifiedAt,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}
}
