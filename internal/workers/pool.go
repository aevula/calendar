package workers

import (
	"context"
	"sync"

	"github.com/aevula/interview-hustlers-calendar/internal/queues"
)

type Pool interface {
	Start()
	Stop(ctx context.Context)
}

type pool struct {
	queue   queues.Queue
	workers []Worker
	wg      sync.WaitGroup
}

func NewPool(size int, queue queues.Queue, opts WorkerOptions) Pool {
	workers := make([]Worker, 0, size)
	for range size {
		workers = append(workers, NewWorker(queue, opts))
	}

	return &pool{queue: queue, workers: workers}
}

func (p *pool) Start() {
	for i := range p.workers {
		p.wg.Go(p.workers[i].Start)
	}
}

func (p *pool) Stop(ctx context.Context) {
	defer p.wg.Wait()

	for i := range p.workers {
		p.workers[i].Stop(ctx)
	}
}
