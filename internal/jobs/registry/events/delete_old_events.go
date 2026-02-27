package events

import (
	tasksDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/tasks"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs"
	eventsJobs "github.com/aevula/interview-hustlers-calendar/internal/jobs/events"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs/registry"
)

type deleteOldEventsFactory struct {
	repo eventsJobs.DeleteOldEventsRepository
}

func NewDeleteOldEvents(repo eventsJobs.DeleteOldEventsRepository) registry.JobFactory {
	return &deleteOldEventsFactory{repo: repo}
}

func (fac *deleteOldEventsFactory) Create(tasksDomain.Task) (jobs.Job, error) {
	return eventsJobs.NewDeleteOldEvents(fac.repo), nil
}
