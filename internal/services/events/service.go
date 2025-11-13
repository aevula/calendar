package events

import (
	"context"
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/repository"
)

type CreateEvent struct {
	Title        string
	Description  string
	UserId       int
	StartAt      time.Time
	Duration     time.Duration
	NotifyOffset time.Duration
}

type UpdateEvent struct {
	ID           int
	Title        string
	Description  string
	UserId       int
	StartAt      time.Time
	Duration     time.Duration
	NotifyOffset time.Duration
	NotifiedAt   time.Time
}

type DeleteEvent struct {
	ID int
}

type ListEvents struct{}

func (e *CreateEvent) toRepo() repository.Event {
	return repository.Event{
		Title:        e.Title,
		Description:  e.Description,
		UserId:       e.UserId,
		StartAt:      e.StartAt,
		Duration:     e.Duration,
		NotifyOffset: e.NotifyOffset,
	}
}

func (e *UpdateEvent) toRepo() repository.Event {
	return repository.Event{
		ID:           e.ID,
		Title:        e.Title,
		Description:  e.Description,
		UserId:       e.UserId,
		StartAt:      e.StartAt,
		Duration:     e.Duration,
		NotifyOffset: e.NotifyOffset,
		NotifiedAt:   e.NotifiedAt,
	}
}

func (e *DeleteEvent) toRepo() repository.Event {
	return repository.Event{
		ID: e.ID,
	}
}

func Create(ctx context.Context, rep repository.EventRepository, event CreateEvent) (repository.Event, error) {
	return rep.Create(ctx, event.toRepo())
}

func Update(ctx context.Context, rep repository.EventRepository, event UpdateEvent) (repository.Event, error) {
	return rep.Update(ctx, event.toRepo())
}

func Delete(ctx context.Context, rep repository.EventRepository, event DeleteEvent) error {
	return rep.Delete(ctx, event.toRepo())
}

func List(ctx context.Context, rep repository.EventRepository) ([]repository.Event, error) {
	return rep.List(ctx)
}
