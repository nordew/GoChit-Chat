package dao

import (
	"user/internal/model"
	repoModel "user/internal/repository/user/model"
)

func ToUserFromRepo(user *repoModel.User) *model.User {
	return &model.User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		RefreshToken: user.RefreshToken,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
