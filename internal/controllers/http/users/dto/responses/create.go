package responses

import (
	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/users"
)

type CreateUserResponse struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

func ToCreateUserResponse(user domain.User) CreateUserResponse {
	return CreateUserResponse{
		ID:    int(user.ID),
		Login: user.Login,
		Name:  user.Name,
	}
}
