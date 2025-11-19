package users

import (
	"context"

	"github.com/aevula/interview-hustlers-calendar/internal/repository"
)

type UsersService interface {
	Create(ctx context.Context, cmd CreateUserCommand) (repository.User, error)
}

type usersService struct {
	repo repository.UserRepository
}

func NewUsersService(repo repository.UserRepository) UsersService {
	return &usersService{repo: repo}
}
