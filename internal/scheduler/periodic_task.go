package scheduler

import (
	"time"

	tasksService "github.com/aevula/interview-hustlers-calendar/internal/services/tasks"
)

type PeriodicTask struct {
	Interval time.Duration
	Name     string
	Queue    string
	Payload  map[string]any
	StartAt  *time.Time
}

func (t *PeriodicTask) ToCreateTaskCommand() tasksService.CreateTaskCommand {
	return tasksService.CreateTaskCommand{
		Name:    t.Name,
		Queue:   t.Queue,
		Payload: t.Payload,
		StartAt: t.StartAt,
	}
}
