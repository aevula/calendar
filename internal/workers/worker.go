package workers

import (
	"context"

	"github.com/aevula/interview-hustlers-calendar/internal/jobs"
)

type Worker interface {
	Add(job jobs.Job)
	Run(ctx context.Context)
	Stop(ctx context.Context) error
}

type worker struct {
	available chan<- Worker
	jobs      chan jobs.Job
}

func New(available chan<- Worker) Worker {
	return &worker{
		available: available,
		jobs:      make(chan jobs.Job, 1),
	}
}

func (w *worker) Add(job jobs.Job) {
	w.jobs <- job
}

func (w *worker) Run(ctx context.Context) {

	go func() {
		for {
			w.available <- w

			select {
			case job := <-w.jobs:
				job.Run(ctx)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (w *worker) Stop(ctx context.Context) error {
	return nil
}
