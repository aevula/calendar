package registry

import (
	tasksDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/tasks"
)

type JobFactoryRegistry interface {
	Register(name tasksDomain.TaskName, factory JobFactory)
	Get(name tasksDomain.TaskName) (JobFactory, bool)
}

type jobFactoryRegistry struct {
	factories map[tasksDomain.TaskName]JobFactory
}

func NewJobFactoryRegistry() JobFactoryRegistry {
	return &jobFactoryRegistry{factories: make(map[tasksDomain.TaskName]JobFactory)}
}

func (jfr *jobFactoryRegistry) Register(taskName tasksDomain.TaskName, factory JobFactory) {
	jfr.factories[taskName] = factory
}

func (jfr *jobFactoryRegistry) Get(taskName tasksDomain.TaskName) (JobFactory, bool) {
	f, ok := jfr.factories[taskName]
	return f, ok
}
