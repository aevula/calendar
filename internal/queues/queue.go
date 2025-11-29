package queues

import (
	"context"
	"errors"
	"sync/atomic"

	"github.com/aevula/interview-hustlers-calendar/internal/jobs"
	fifos "github.com/aevula/interview-hustlers-calendar/pkg/queues"
)

var QueueStopped = errors.New("queue stopped")

type Queue interface {
	Name() string
	Start()
	Push(ctx context.Context, job jobs.Job) error
	Pop(ctx context.Context) (jobs.Job, error)
	Stop(ctx context.Context)
}

type queue struct {
	fifo  fifos.Queue[jobs.Job]
	name  string
	state atomic.Bool
}

const FixedQueueCap = 1_000

func NewQueue(name string) Queue {
	return &queue{
		fifo:  fifos.NewFixedQueue[jobs.Job](FixedQueueCap),
		name:  name,
		state: atomic.Bool{},
	}
}

func (q *queue) Name() string {
	return q.name
}

func (q *queue) Start() {
	q.state.Store(true)
}

func (q *queue) Push(ctx context.Context, job jobs.Job) error {
	return q.fifo.Enqueue(ctx, job)
}

func (q *queue) Pop(ctx context.Context) (jobs.Job, error) {
	if q.state.Load() {
		return q.fifo.Dequeue(ctx)
	}

	return nil, QueueStopped
}

func (q *queue) Stop(context.Context) {
	q.state.Store(false)
}
