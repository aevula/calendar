package db

import "context"

type Db interface {
	Ping(ctx context.Context) error
	Exec(ctx context.Context, sql string, args ...any) error
	Close(ctx context.Context) error
	// BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	// Close(ctx context.Context) error
	// Config() *pgx.ConnConfig
	// CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	// Deallocate(ctx context.Context, name string) error
	// DeallocateAll(ctx context.Context) error
	// Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	// IsClosed() bool
	// LoadType(ctx context.Context, typeName string) (*pgtype.Type, error)
	// LoadTypes(ctx context.Context, typeNames []string) ([]*pgtype.Type, error)
	// PgConn() *pgconn.PgConn
	// Prepare(ctx context.Context, name string, sql string) (sd *pgconn.StatementDescription, err error)
	// Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	// QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	// SendBatch(ctx context.Context, b *pgx.Batch) (br pgx.BatchResults)
	// TypeMap() *pgtype.Map
	// WaitForNotification(ctx context.Context) (*pgconn.Notification, error)
}

func a() {
}
