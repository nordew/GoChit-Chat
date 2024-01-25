package model

import (
	"github.com/google/uuid"
	"time"
)

type UserRole string

const (
	Admin   UserRole = "ADMIN"
	Default UserRole = "User"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	Password     string
	RefreshToken string
	Role         UserRole
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
