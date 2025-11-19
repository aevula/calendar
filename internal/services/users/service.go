package users

import (
	"context"

	"github.com/aevula/interview-hustlers-calendar/internal/repository"
)

type UsersService interface {
	Create(ctx context.Context, cmd CreateUserCommand) (repository.User, error)
}

type usersService struct {
	repo UserRepository
}

type UserRepository interface {
	Create(context.Context, repository.User) (repository.User, error)
}

func NewUsersService(repo UserRepository) UsersService {
	return &usersService{repo: repo}
}
