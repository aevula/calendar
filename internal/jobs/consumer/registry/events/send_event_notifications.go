package events

import (
	"github.com/aevula/interview-hustlers-calendar/internal/infra/kafka"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs/consumer/events"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs/consumer/registry"
)

type sendEventNotificationFactory struct {
	repo events.SendEventNotificationRepository
}

func NewSendEventNotification(repo events.SendEventNotificationRepository) registry.JobFactory {
	return &sendEventNotificationFactory{repo: repo}
}

func (fac *sendEventNotificationFactory) Create(record kafka.Record) (jobs.Job, error) {
	return events.NewSendEventNotification(fac.repo, record.Value), nil
}
