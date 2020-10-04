package models

import (
	"fmt"
	"time"
)

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  string    `json:"password"`
	DeletedAt time.Time `pg:",soft_delete"`
}

func (u User) String() string {
	return fmt.Sprintf("%d %s %v %v", u.ID, u.Email, u.CreatedAt, u.DeletedAt)
}