package registry

import "github.com/aevula/interview-hustlers-calendar/internal/infra/kafka"

type JobFactoryRegistry interface {
	Register(name kafka.Topic, factory JobFactory)
	Get(name kafka.Topic) (JobFactory, bool)
}

type jobFactoryRegistry struct {
	factories map[kafka.Topic]JobFactory
}

func NewJobFactoryRegistry() JobFactoryRegistry {
	return &jobFactoryRegistry{factories: make(map[kafka.Topic]JobFactory)}
}

func (jfr *jobFactoryRegistry) Register(topic kafka.Topic, factory JobFactory) {
	jfr.factories[topic] = factory
}

func (jfr *jobFactoryRegistry) Get(topic kafka.Topic) (JobFactory, bool) {
	f, ok := jfr.factories[topic]
	return f, ok
}
