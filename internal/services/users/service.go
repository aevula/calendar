package users

import (
	"context"

	usersDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/users"
)

type UsersService interface {
	Create(ctx context.Context, cmd CreateUserCommand) (usersDomain.User, error)
}

type usersService struct {
	repo UserRepository
}

type UserRepository interface {
	Create(context.Context, usersDomain.User) (usersDomain.User, error)
}

func NewUsersService(repo UserRepository) UsersService {
	return &usersService{repo: repo}
}
