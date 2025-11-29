package events

import (
	"context"
	"errors"
	"time"

	eventsDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
	"github.com/aevula/interview-hustlers-calendar/internal/infra/kafka"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs"
)

type SendEventsNotificationsRepository interface {
	FindAllNotNotifiedByNotifyAt(ctx context.Context, filter eventsDomain.NotifyAtFilter) ([]eventsDomain.Event, error)
}

type sendEventsNotifications struct {
	eventsRepo SendEventsNotificationsRepository
	producer   kafka.Producer
}

func NewSendEventsNotifications(eventsRepo SendEventsNotificationsRepository, producer kafka.Producer) jobs.Job {
	return &sendEventsNotifications{
		eventsRepo: eventsRepo,
		producer:   producer,
	}
}

func (job *sendEventsNotifications) Run(ctx context.Context) error {
	filter := eventsDomain.NotifyAtFilter{LTE: time.Now()}

	events, err := job.eventsRepo.FindAllNotNotifiedByNotifyAt(ctx, filter)
	if err != nil {
		return err
	}

	for i := range events {
		data := map[string]any{"eventID": events[i].ID}

		if err = job.producer.ProduceSync(ctx, "event_notifications", data); err != nil {
			err = errors.Join(err)
		}
	}

	return err
}
