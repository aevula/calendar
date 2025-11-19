package users

import "time"

type UserID int

type User struct {
	ID        UserID
	Login     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
