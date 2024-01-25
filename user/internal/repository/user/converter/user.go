package converter

import (
	"user/internal/model"
	repoModel "user/internal/repository/user/model"
)

func toUserFromRepo(user *repoModel.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
