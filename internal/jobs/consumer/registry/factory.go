package registry

import (
	"github.com/aevula/interview-hustlers-calendar/internal/infra/kafka"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs"
)

type JobFactory interface {
	Create(record kafka.Record) (jobs.Job, error)
}
