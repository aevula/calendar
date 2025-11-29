package jobs

import "context"

type Job interface {
	Run(ctx context.Context) error
	// Priority() int // TODO
}
