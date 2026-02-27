package users

import (
	"context"

	usersDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/users"
)

type CreateUserCommand struct {
	Login string
	Name  string
}

func (s *usersService) Create(ctx context.Context, cmd CreateUserCommand) (usersDomain.User, error) {
	user := usersDomain.User{
		Login: cmd.Login,
		Name:  cmd.Name,
	}
	return s.repo.Create(ctx, user)
}
