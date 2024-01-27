package dao

import (
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"google.golang.org/protobuf/types/known/timestamppb"
	"user/internal/model"
)

func ToUserFromCreateRequest(req *desc.CreateUserRequest) *model.User {
	return &model.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func ToUserFromUpdateRequest(req *desc.UpdateUserRequest) *model.UserUpdate {
	return &model.UserUpdate{
		ID:          req.GetId(),
		Name:        req.GetName().GetValue(),
		Email:       req.GetEmail(),
		Password:    req.GetPassword(),
		OldPassword: req.GetOldPassword().GetValue(),
		NewPassword: req.GetNewPassword().GetValue(),
	}
}

func ToResponseFromUser(user *model.User) *desc.GetUserResponse {
	createdAtTimestamp := timestamppb.New(user.CreatedAt)
	updatedAtTimestamp := timestamppb.New(user.UpdatedAt)

	return &desc.GetUserResponse{
		Id:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      int32(user.Role),
		CreatedAt: createdAtTimestamp,
		UpdatedAt: updatedAtTimestamp,
	}
}
