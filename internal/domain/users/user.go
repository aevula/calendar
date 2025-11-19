package users

import "time"

type User struct {
	ID        int
	Login     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
