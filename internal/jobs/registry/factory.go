package registry

import (
	tasksDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/tasks"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs"
)

type JobFactory interface {
	Create(task tasksDomain.Task) (jobs.Job, error)
}
