package postgres

import (
	"context"
	"fmt"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"github.com/aevula/interview-hustlers-calendar/internal/databases"
	"github.com/jackc/pgx/v5"
)

type DB struct {
	conn *pgx.Conn
}

func New(ctx context.Context, cfg config.Config) (databases.DB, error) {
	con, err := pgx.Connect(ctx, connString(cfg))
	if err != nil {
		return nil, err
	}
	return &DB{conn: con}, nil
}

func connString(cfg config.Config) string {
	return fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
}

func (db *DB) Ping(ctx context.Context) error {
	return db.conn.Ping(ctx)
}

func (db *DB) Close(ctx context.Context) error {
	return db.conn.Close(ctx)
}

func (db *DB) QueryRow(ctx context.Context, sql string, args ...any) (databases.Row, error) {
	row := db.conn.QueryRow(ctx, sql, args...)
	return databases.Row(row), nil
}

func (db *DB) Query(ctx context.Context, sql string, args ...any) (databases.Rows, error) {
	rows, err := db.conn.Query(ctx, sql, args...)
	return databases.Rows(rows), err
}

func (db *DB) Exec(ctx context.Context, sql string, args ...any) error {
	_, err := db.conn.Exec(ctx, sql, args...)
	return err
}
