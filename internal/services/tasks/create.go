package tasks

import (
	"context"
	"time"

	tasksDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/tasks"
)

type CreateTaskCommand struct {
	Name    string
	Queue   string
	Payload map[string]any
	StartAt *time.Time
}

func (s *tasksService) Create(ctx context.Context, cmd CreateTaskCommand) (tasksDomain.Task, error) {
	startAt := time.Now()
	if cmd.StartAt != nil {
		startAt = *cmd.StartAt
	}

	task := tasksDomain.Task{
		Name:    tasksDomain.TaskName(cmd.Name),
		Queue:   tasksDomain.QueueName(cmd.Queue),
		Payload: tasksDomain.TaskPayload(cmd.Payload),
		StartAt: startAt,
	}
	return s.repo.Create(ctx, task)
}
