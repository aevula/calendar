package requests

import (
	"github.com/aevula/interview-hustlers-calendar/internal/services/users"
)

type CreateUserRequest struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

func (r CreateUserRequest) ToCommand() users.CreateUserCommand {
	return users.CreateUserCommand{
		Login: r.Login,
		Name:  r.Name,
	}
}
