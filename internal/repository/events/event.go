package events

import (
	"time"

	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
)

type Event struct {
	ID          int
	Title       string
	Description string
	UserId      int
	StartAt     time.Time
	Duration    time.Duration
	NotifyAt    time.Time
	NotifiedAt  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func FromDomain(event domain.Event) Event {
	return Event{
		ID:          int(event.ID),
		Title:       event.Title,
		Description: event.Description,
		UserId:      int(event.UserId),
		StartAt:     event.StartAt,
		Duration:    event.Duration,
		NotifyAt:    event.NotifyAt,
		NotifiedAt:  event.NotifiedAt,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}
}

func (event Event) ToDomain() domain.Event {
	return domain.Event{
		ID:          domain.EventID(event.ID),
		Title:       event.Title,
		Description: event.Description,
		UserId:      domain.UserID(event.UserId),
		StartAt:     event.StartAt,
		Duration:    event.Duration,
		NotifyAt:    event.NotifyAt,
		NotifiedAt:  event.NotifiedAt,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}
}
