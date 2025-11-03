package postgres

import (
	"context"
	"fmt"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"github.com/jackc/pgx/v5"
)

type Db struct {
	conn *pgx.Conn
}

func New(ctx context.Context, cfg config.Config) (*Db, error) {
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
