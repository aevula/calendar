package users

import (
	"context"

	"github.com/aevula/interview-hustlers-calendar/internal/repository"
)

type CreateUser struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

func (e *CreateUser) toRepo() repository.User {
	return repository.User{
		Login: e.Login,
		Name:  e.Name,
	}
}

func Create(ctx context.Context, rep repository.UserRepository, user CreateUser) (repository.User, error) {
	return rep.Create(ctx, user.toRepo())
}
