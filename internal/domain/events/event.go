package events

import (
	"time"
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

func (e *Event) EnsureNotifyAt() {
	if e.NotifyAt.IsZero() {
		e.NotifyAt = e.StartAt
	}
}
