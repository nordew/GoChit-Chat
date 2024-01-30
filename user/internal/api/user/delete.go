package user

import (
	"context"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	err := i.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error deleting user")
	}

	return &emptypb.Empty{}, nil
}
