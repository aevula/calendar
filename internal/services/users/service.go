package users

import (
	"context"

	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/users"
)

type UsersService interface {
	Create(ctx context.Context, cmd CreateUserCommand) (domain.User, error)
}

type usersService struct {
	repo UserRepository
}

type UserRepository interface {
	Create(context.Context, domain.User) (domain.User, error)
}

func NewUsersService(repo UserRepository) UsersService {
	return &usersService{repo: repo}
}
