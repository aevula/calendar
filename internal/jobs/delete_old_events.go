package jobs

import (
	"context"
	"sync"
	"time"

	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
	"github.com/aevula/interview-hustlers-calendar/internal/logging"
)

type deleteOldEvents struct {
	eventsRepo EventRepository
	logger     logging.Logger
}

type EventRepository interface {
	Delete(ctx context.Context, id domain.EventID) error
	FindAllByNotifiedAt(ctx context.Context, filter domain.NotifiedAtFilter) ([]domain.Event, error)
}

func NewDeleteOldEvents(eventsRepo EventRepository, logger logging.Logger) Job {
	return &deleteOldEvents{
		eventsRepo: eventsRepo,
		logger:     logger.With(logger.String("tag", "delete_old_events")),
	}
}

func (job *deleteOldEvents) Run(ctx context.Context) {
	filter := domain.NotifiedAtFilter{LTE: time.Now().AddDate(1, 0, 0)}

	events, err := job.eventsRepo.FindAllByNotifiedAt(ctx, filter)
	if err != nil {
		job.logger.Error(err.Error())
		return
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()

	for i := range events {
		wg.Go(func() {
			if err := job.eventsRepo.Delete(ctx, events[i].ID); err != nil {
				job.logger.Error(err.Error())
			} else {
				job.logger.Debug("Deleted", job.logger.Int("id", int(events[i].ID)))
			}
		})
	}
}
