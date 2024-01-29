package dao

import (
	"user/internal/model"
	"user/internal/service"

	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromCreateRequest(req *desc.CreateUserRequest) *model.User {
	return &model.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func ToUserFromUpdateRequest(req *desc.UpdateUserRequest) *service.UpdateUserRequest {
	user := &model.UserUpdate{
		ID:          req.GetId(),
		Email:       req.GetEmail().GetValue(),
		Name:        req.GetName().GetValue(),
		Password:    req.GetPassword(),
		OldPassword: req.GetOldPassword().GetValue(),
		NewPassword: req.NewPassword.GetValue(),
	}

	return &service.UpdateUserRequest{
		User: user,
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

func ToResponseFromCreateUserResponse(resp *service.CreateUserResponse) *desc.CreateUserResponse {
	return &desc.CreateUserResponse{
		Id:           resp.Id,
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}
}
