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
	Ping(ctx context.Context) error
	Close(ctx context.Context) error
	Exec(ctx context.Context, sql string, args ...any) error
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) (Row, error)
}
