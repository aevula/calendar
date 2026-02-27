package queues

import "context"

type FixedQueue[T any] struct {
	fifo chan T
}

func NewFixedQueue[T any](capacity int) Queue[T] {
	return &FixedQueue[T]{fifo: make(chan T, capacity)}
}

func (q *FixedQueue[T]) Enqueue(ctx context.Context, val T) (err error) {
	select {
	case q.fifo <- val:
	case <-ctx.Done():
		err = ctx.Err()
	}
	return
}

func (q *FixedQueue[T]) Dequeue(ctx context.Context) (val T, err error) {
	select {
	case val = <-q.fifo:
	case <-ctx.Done():
		err = ctx.Err()
	}
	return
}
