package events

import (
	"context"
	"time"

	eventsDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs"
)

type DeleteOldEventsRepository interface {
	Delete(ctx context.Context, id eventsDomain.EventID) error
	DeleteAllByNotifiedAt(ctx context.Context, filter eventsDomain.NotifiedAtFilter) ([]eventsDomain.Event, error)
}

type deleteOldEvents struct {
	eventsRepo DeleteOldEventsRepository
}

func NewDeleteOldEvents(eventsRepo DeleteOldEventsRepository) jobs.Job {
	return &deleteOldEvents{
		eventsRepo: eventsRepo,
	}
}

func (job *deleteOldEvents) Run(ctx context.Context) error {
	filter := eventsDomain.NotifiedAtFilter{LTE: time.Now().AddDate(1, 0, 0)}

	_, err := job.eventsRepo.DeleteAllByNotifiedAt(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
