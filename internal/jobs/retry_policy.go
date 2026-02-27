package jobs

import "time"

type RetryPolicy interface {
	ShouldRetry(err error, attempt int) (bool, time.Duration)
}
