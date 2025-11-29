package tasks

import (
	"context"

	tasksDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/tasks"
)

type TasksService interface {
	Create(ctx context.Context, cmd CreateTaskCommand) (tasksDomain.Task, error)
}

type tasksService struct {
	repo TaskRepository
}

type TaskRepository interface {
	Create(context.Context, tasksDomain.Task) (tasksDomain.Task, error)
}

func NewTasksService(repo TaskRepository) TasksService {
	return &tasksService{repo: repo}
}
