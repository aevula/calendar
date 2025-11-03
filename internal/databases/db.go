package databases

import (
	"context"
)

type Db interface {
	Close(ctx context.Context) error
}
