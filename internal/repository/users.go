package repository

import (
	"context"

	"github.com/aevula/interview-hustlers-calendar/internal/databases"
	usersDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/users"
	usersRepo "github.com/aevula/interview-hustlers-calendar/internal/repository/users"
)

type UserRepository interface {
	Create(context.Context, usersDomain.User) (usersDomain.User, error)
}

type userRepository struct {
	db databases.DB
}

func NewUserRepository(db databases.DB) UserRepository {
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

func (r *userRepository) Create(ctx context.Context, user usersDomain.User) (usersDomain.User, error) {
	rUser := usersRepo.FromDomain(user)
	zero := usersRepo.User{}

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
