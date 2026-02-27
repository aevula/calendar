package tasks

import (
	"time"

	tasksDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/tasks"
)

type Task struct {
	ID        int
	Name      string
	Queue     string
	Payload   map[string]any
	Status    string
	Attempt   int
	StartAt   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDomain(task tasksDomain.Task) Task {
	return Task{
		ID:        int(task.ID),
		Name:      string(task.Name),
		Queue:     string(task.Queue),
		Payload:   map[string]any(task.Payload),
		Status:    string(task.Status),
		Attempt:   task.Attempt,
		StartAt:   task.StartAt,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

func (task Task) ToDomain() tasksDomain.Task {
	return tasksDomain.Task{
		ID:        tasksDomain.TaskID(task.ID),
		Name:      tasksDomain.TaskName(task.Name),
		Queue:     tasksDomain.QueueName(task.Queue),
		Payload:   map[string]any(task.Payload),
		Status:    tasksDomain.TaskStatus(task.Status),
		Attempt:   task.Attempt,
		StartAt:   task.StartAt,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}
