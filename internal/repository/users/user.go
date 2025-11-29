package users

import (
	"time"

	usersDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/users"
)

type User struct {
	ID        int
	Login     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDomain(user usersDomain.User) User {
	return User{
		ID:        int(user.ID),
		Login:     user.Login,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (user User) ToDomain() usersDomain.User {
	return usersDomain.User{
		ID:        usersDomain.UserID(user.ID),
		Login:     user.Login,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
