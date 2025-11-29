package requests

import (
	usersService "github.com/aevula/interview-hustlers-calendar/internal/services/users"
)

type CreateUserRequest struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

func (r CreateUserRequest) ToCommand() usersService.CreateUserCommand {
	return usersService.CreateUserCommand{
		Login: r.Login,
		Name:  r.Name,
	}
}
