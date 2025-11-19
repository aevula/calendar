package jobs

import (
	"context"
	"sync"
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/logging"
	"github.com/aevula/interview-hustlers-calendar/internal/repository"
)

type notifyEvent struct {
	eventsRepo EventRepository
	logger     logging.Logger
}

type EventRepository interface {
	Update(ctx context.Context, event repository.Event) (repository.Event, error)
	AllNotifyable(ctx context.Context, from time.Time) ([]repository.Event, error)
}

func NewNotifyEvent(eventsRepo EventRepository, logger logging.Logger) Job {
	return &notifyEvent{
		eventsRepo: eventsRepo,
		logger:     logger.With(logger.String("tag", "notify_event")),
	}
}

func (job *notifyEvent) Run(ctx context.Context) {
	events, err := job.eventsRepo.AllNotifyable(ctx, time.Now())
	if err != nil {
		job.logger.Error(err.Error())
		return
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()

	for i := range events {
		wg.Go(func() {
			events[i].NotifiedAt = time.Now()

			if _, err := job.eventsRepo.Update(ctx, events[i]); err != nil {
				job.logger.Error(err.Error())
			} else {
				job.logger.Debug("Notified", job.logger.Int("id", events[i].ID))
			}
		})
	}
}
