package workers

import (
	"context"
	"errors"
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/jobs"
	"github.com/aevula/interview-hustlers-calendar/internal/queues"
)

type Worker interface {
	Start()
	Stop(ctx context.Context)
}

type WorkerOptions struct {
	Retry             jobs.RetryPolicy
	QueueStoppedSleep time.Duration
}

type worker struct {
	queue queues.Queue
	opts  WorkerOptions
	ctx   context.Context
	stop  context.CancelFunc
}

func NewWorker(queue queues.Queue, opts WorkerOptions) Worker {
	ctx, stop := context.WithCancel(context.Background())
	return &worker{queue: queue, opts: opts, ctx: ctx, stop: stop}
}

func (w *worker) Start() {
	for {
		select {
		case <-w.ctx.Done():
			return
		default:
		}

		job, err := w.queue.Pop(w.ctx)
		if err != nil && errors.Is(err, queues.QueueStopped) {
			<-time.After(w.opts.QueueStoppedSleep)
			continue
		}

		w.run(job)
	}
}

func (w *worker) Stop(ctx context.Context) {
	w.stop()
}

func (w *worker) run(job jobs.Job) {
	attempt := 1
	for {
		select {
		case <-w.ctx.Done():
			return
		default:
		}

		err := job.Run(w.ctx)
		if err == nil {
			return
		}

		retry, sleep := w.opts.Retry.ShouldRetry(err, attempt)
		if !retry {
			return
		}

		select {
		case <-w.ctx.Done():
			return
		case <-time.After(sleep):
		}

		attempt++
	}
}
