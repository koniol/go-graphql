package models

import (
	"time"
)

type Document struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	UserId    int        `json:"user_id"`
	DeletedAt *time.Time `pg:",soft_delete"`
}
