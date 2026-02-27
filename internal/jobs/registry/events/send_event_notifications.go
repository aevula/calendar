package events

import (
	tasksDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/tasks"
	"github.com/aevula/interview-hustlers-calendar/internal/infra/kafka"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs"
	eventsJobs "github.com/aevula/interview-hustlers-calendar/internal/jobs/events"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs/registry"
)

type sendEventNotificationsFactory struct {
	repo     eventsJobs.SendEventsNotificationsRepository
	producer kafka.Producer
}

func NewSendEventsNotifications(repo eventsJobs.SendEventsNotificationsRepository, producer kafka.Producer) registry.JobFactory {
	return &sendEventNotificationsFactory{repo: repo, producer: producer}
}

func (fac *sendEventNotificationsFactory) Create(tasksDomain.Task) (jobs.Job, error) {
	return eventsJobs.NewSendEventsNotifications(fac.repo, fac.producer), nil
}
