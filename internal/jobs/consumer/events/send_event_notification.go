package events

import (
	"context"
	"fmt"

	eventsDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
	"github.com/aevula/interview-hustlers-calendar/internal/infra/kafka"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs"
)

type SendEventNotificationRepository interface {
	FindById(ctx context.Context, id eventsDomain.EventID) (eventsDomain.Event, error)
}

type sendEventNotification struct {
	repo SendEventNotificationRepository
	data kafka.RecordValue
}

func NewSendEventNotification(repo SendEventNotificationRepository, data kafka.RecordValue) jobs.Job {
	return &sendEventNotification{repo: repo, data: data}
}

func (job *sendEventNotification) Run(ctx context.Context) error {
	eventID, ok := job.data["eventID"]
	if !ok {
		return fmt.Errorf("empty 'eventID'")
	}

	id, ok := eventID.(int)
	if !ok {
		return fmt.Errorf("invalid 'eventID'")
	}

	event, err := job.repo.FindById(ctx, eventsDomain.EventID(id))
	if err != nil {
		return err
	}

	_, err = fmt.Printf("Event %v at %v for %v:\n%v", event.Title, event.StartAt, event.Duration, event.Description)
	return err
}
