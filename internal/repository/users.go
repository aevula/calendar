package repository

import (
	"context"
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/databases"
)

type User struct {
	ID        int
	Login     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	Create(context.Context, User) (User, error)
}

type userRepository struct {
	db databases.Db
}

func NewUserRepository(db databases.Db) UserRepository {
	return &userRepository{db: db}
}

const userCreateSQL = `
INSERT INTO users (
	login, name
) VALUES (
	$1,	   $2
)
RETURNING
	id, login, name, created_at, updated_at
`

func (r *userRepository) Create(ctx context.Context, user User) (User, error) {
	empty := User{}

	row, err := r.db.QueryRow(ctx, userCreateSQL,
		user.Login,
		user.Name,
	)

	if err != nil {
		return empty, err
	}

	err = row.Scan(
		&empty.ID,
		&empty.Login,
		&empty.Name,
		&empty.CreatedAt,
		&empty.UpdatedAt,
	)

	return empty, err
}
