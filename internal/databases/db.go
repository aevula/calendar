package databases

import (
	"context"
)

type Row interface {
	Scan(dest ...any) error
}

type Rows interface {
	Scan(dest ...any) error
	Next() bool // Todo: #CollectRows() or #ForEachRow()
	Err() error
	Close()
}

type DB interface {
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) (Row, error)
	Close(ctx context.Context) error
	Ping(ctx context.Context) error
}
