package postgres

import (
	"context"
	"fmt"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"github.com/aevula/interview-hustlers-calendar/internal/databases"
	"github.com/jackc/pgx/v5"
)

type Db struct {
	conn *pgx.Conn
}

func New(ctx context.Context, cfg config.Config) (databases.Db, error) {
	con, err := pgx.Connect(ctx, connString(cfg))
	if err != nil {
		return nil, err
	}
	return &Db{conn: con}, nil
}

func connString(cfg config.Config) string {
	return fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v",
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Name,
	)
}

func (db *Db) Close(ctx context.Context) error {
	return db.conn.Close(ctx)
}

func (db *Db) QueryRow(ctx context.Context, sql string, args ...any) (databases.Row, error) {
	row := db.conn.QueryRow(ctx, sql, args...)
	return databases.Row(row), nil
}

func (db *Db) Query(ctx context.Context, sql string, args ...any) (databases.Rows, error) {
	rows, err := db.conn.Query(ctx, sql, args...)
	return databases.Rows(rows), err
}

func (db *Db) Exec(ctx context.Context, sql string, args ...any) error {
	_, err := db.conn.Exec(ctx, sql, args...)
	return err
}
