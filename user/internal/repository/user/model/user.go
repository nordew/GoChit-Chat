package model

import (
	"time"
)

type Role string

type User struct {
	ID           string    `db:"id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	RefreshToken string    `db:"refresh_token"`
	Role         int       `db:"role"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
