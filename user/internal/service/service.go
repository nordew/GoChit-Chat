package service

import (
	"context"
	"user/internal/model"
	userErrors "user/pkg/user_errors"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) (*CreateUserResponse, *userErrors.CustomErr)
	Get(ctx context.Context, email string) (*model.User, *userErrors.CustomErr)
	Update(ctx context.Context, req *UpdateUserRequest) *userErrors.CustomErr
	Delete(ctx context.Context, id string) *userErrors.CustomErr
	Login(ctx context.Context, email string, password string) (string, string, *userErrors.CustomErr)
	ParseAccessToken(ctx context.Context, token string) (string, string, *userErrors.CustomErr)
	Refresh(ctx context.Context, token string) (string, string, *userErrors.CustomErr)
}

type CreateUserResponse struct {
	Id           string
	AccessToken  string
	RefreshToken string
}

type UpdateUserRequest struct {
	User *model.UserUpdate
}
