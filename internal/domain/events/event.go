package events

import (
	"time"

	usersDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/users"
)

type EventID int
type UserID usersDomain.UserID

type Event struct {
	ID          EventID
	Title       string
	Description string
	UserId      UserID
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
