package repository

import (
	"context"

	"github.com/aevula/interview-hustlers-calendar/internal/databases"
	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/users"
	repo "github.com/aevula/interview-hustlers-calendar/internal/repository/users"
)

type UserRepository interface {
	Create(context.Context, domain.User) (domain.User, error)
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

func (r *userRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	rUser := repo.FromDomain(user)
	zero := repo.User{}

	row, err := r.db.QueryRow(ctx, userCreateSQL,
		rUser.Login,
		rUser.Name,
	)

	if err != nil {
		return zero.ToDomain(), err
	}

	err = row.Scan(
		&zero.ID,
		&zero.Login,
		&zero.Name,
		&zero.CreatedAt,
		&zero.UpdatedAt,
	)

	return zero.ToDomain(), err
}
