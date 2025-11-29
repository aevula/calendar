package dispatcher

import (
	"context"
	"fmt"
	"time"

	tasksDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/tasks"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs/registry"
	"github.com/aevula/interview-hustlers-calendar/internal/queues"
	"github.com/aevula/interview-hustlers-calendar/internal/repository"
)

type JobsDispatcher interface {
	Start()
	Stop(ctx context.Context)
}

type DispatcherOptions struct {
	RefreshRate time.Duration
}

type jobsDispatcher struct {
	queues       map[string]queues.Queue
	jobsRegistry registry.JobFactoryRegistry
	tasksRepo    repository.TaskRepository
	opts         DispatcherOptions
	ctx          context.Context
	stop         context.CancelFunc
}

func NewJobsDispatcher(qs []queues.Queue, reg registry.JobFactoryRegistry, repo repository.TaskRepository, opts DispatcherOptions) JobsDispatcher {
	ctx, stop := context.WithCancel(context.Background())

	qMap := make(map[string]queues.Queue, len(qs))
	for _, q := range qs {
		qMap[q.Name()] = q
	}

	return &jobsDispatcher{
		queues:       qMap,
		jobsRegistry: reg,
		tasksRepo:    repo,
		opts:         opts,
		ctx:          ctx,
		stop:         stop,
	}
}

func (jd *jobsDispatcher) Start() {
	go func() {
		for {
			select {
			case <-jd.ctx.Done():
				return
			case <-time.After(jd.opts.RefreshRate):
				_ = jd.dispatch()
			}
		}
	}()
}

func (jd *jobsDispatcher) Stop(ctx context.Context) {
	jd.stop()
}

func (jd *jobsDispatcher) dispatch() error {
	tasks, err := jd.tasksRepo.FindAllByStatus(jd.ctx, tasksDomain.StatusFilter{
		Status: tasksDomain.TaskStatusPending,
	})
	if err != nil || len(tasks) == 0 {
		return err
	}

	ids := make([]tasksDomain.TaskID, 0, len(tasks))
	for i := range tasks {
		ids = append(ids, tasks[i].ID)
	}

	if err := jd.tasksRepo.UpdateStatusByIds(jd.ctx, ids, tasksDomain.TaskStatusEnqueued); err != nil {
		return err
	}

	errIds := make([]tasksDomain.TaskID, 0, len(ids))

	for i := range tasks {
		if err := jd.enqueueTask(tasks[i]); err != nil {
			errIds = append(errIds, tasks[i].ID)
		}
	}

	if len(errIds) > 0 {
		return jd.tasksRepo.UpdateStatusByIds(jd.ctx, errIds, tasksDomain.TaskStatusPending)
	}

	return nil
}

func (jd *jobsDispatcher) enqueueTask(task tasksDomain.Task) error {
	queue, ok := jd.queues[string(task.Queue)]
	if !ok {
		return errUnknownQueue(task.Queue)
	}

	factory, ok := jd.jobsRegistry.Get(task.Name)
	if !ok {
		return errUnknownJob(task.Name)
	}

	job, err := factory.Create(task)
	if err != nil {
		return err
	}

	return queue.Push(jd.ctx, job)
}

func errUnknownQueue(q tasksDomain.QueueName) error {
	return fmt.Errorf("unknown queue: %s", q)
}

func errUnknownJob(name tasksDomain.TaskName) error {
	return fmt.Errorf("unknown job: %s", name)
}
