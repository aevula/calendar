package postgres

import (
	"context"
	"fmt"

	"github.com/aevula/interview-hustlers-calendar/internal/application/config"
	"github.com/jackc/pgx/v5"
)

type Db struct {
	*pgx.Conn
}

func New(ctx context.Context, cfg config.Config) (*Db, error) {
	con, err := pgx.Connect(ctx, connString(cfg))
	return &Db{Conn: con}, err
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
