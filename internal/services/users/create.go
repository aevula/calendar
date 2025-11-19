package users

import (
	"context"

	"github.com/aevula/interview-hustlers-calendar/internal/repository"
)

type CreateUserCommand struct {
	Login string
	Name  string
}

func (s *usersService) Create(ctx context.Context, cmd CreateUserCommand) (repository.User, error) {
	user := repository.User{
		Login: cmd.Login,
		Name:  cmd.Name,
	}
	return s.repo.Create(ctx, user)
}
