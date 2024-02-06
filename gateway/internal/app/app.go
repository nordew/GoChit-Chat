package app

import (
	"fmt"
	"gateway/config"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	v1 "gateway/internal/controller/http/v1"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
)

func MustRun() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to run config: %s", err.Error())
	}

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to run config: %s", err.Error())
	}

	userClientConn, err := grpc.Dial("chat-user:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %s", err.Error())
	}

	userClientService := desc.NewUserClient(userClientConn)

	handler := v1.NewHandler(userClientService)

	router := handler.Init()

	log.Printf("starting router at port: %d", cfg.HTTPPort)
	if err := router.Listen(fmt.Sprintf(":%d", cfg.HTTPPort)); err != nil {
		log.Fatalf("failed to run router: %s", err.Error())
	}
}
