package events

import (
	"time"

	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/users"
)

type User struct {
	ID        int
	Login     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDomain(user domain.User) User {
	return User{
		ID:        user.ID,
		Login:     user.Login,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (user User) ToDomain() domain.User {
	return domain.User{
		ID:        user.ID,
		Login:     user.Login,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
