package model

import (
	"github.com/google/uuid"
	"time"
)

type Role string

const (
	AdminRole Role = "ADMIN"
	UserRole  Role = "USER"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	Password     string
	RefreshToken string
	Role         Role
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
