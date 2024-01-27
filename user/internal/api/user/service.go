package user

import (
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"user/internal/service"
)

type Implementation struct {
	desc.UnimplementedUserServer
	userService service.UserService

	log *zap.Logger
}

func Register(server *grpc.Server, userService service.UserService, log *zap.Logger) {
	desc.RegisterUserServer(server, NewUserImplementation(userService, log))
}

func NewUserImplementation(userService service.UserService, log *zap.Logger) *Implementation {
	return &Implementation{
		userService: userService,
		log:         log,
	}
}
