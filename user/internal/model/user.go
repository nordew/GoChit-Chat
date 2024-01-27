package model

import (
	"time"
)

type UserRole string

type User struct {
	ID           string
	Name         string
	Email        string
	Password     string
	RefreshToken string
	Role         int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserUpdate struct {
	ID          string
	Name        string
	Email       string
	Password    bool
	NewPassword string
	OldPassword string
	Role        int
	UpdatedAt   time.Time
}
