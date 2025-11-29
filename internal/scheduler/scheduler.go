package scheduler

import (
	"context"
	"sync"
	"time"

	tasksService "github.com/aevula/interview-hustlers-calendar/internal/services/tasks"
)

type Scheduler interface {
	Start()
	Stop(ctx context.Context)
}

type scheduler struct {
	tasksService tasksService.TasksService
	tasks        []PeriodicTask
	ctx          context.Context
	stop         context.CancelFunc
	wg           sync.WaitGroup
}

func New(tasksService tasksService.TasksService, tasks []PeriodicTask) Scheduler {
	ctx, stop := context.WithCancel(context.Background())
	return &scheduler{tasksService: tasksService, tasks: tasks, ctx: ctx, stop: stop}
}

func (s *scheduler) Start() {
	for i := range s.tasks {
		s.wg.Go(func() {
			s.schedule(s.tasks[i])
		})
	}
}

func (s *scheduler) Stop(ctx context.Context) {
	defer s.wg.Wait()

	s.stop()
}

func (s *scheduler) schedule(task PeriodicTask) {
	for {
		select {
		case <-s.ctx.Done():
			return
		case <-time.After(task.Interval):
		}

		if _, err := s.tasksService.Create(s.ctx, task.ToCreateTaskCommand()); err != nil {
			return
		}
	}
}
