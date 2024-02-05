package app

import (
	"chat/internal/config"
	v1 "chat/internal/controller/http/v1"
	roomRepo "chat/internal/repository/room"
	"chat/internal/service/chat"
	"chat/internal/service/room"
	logger "chat/pkg/logger"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func MustRun() {
	logger := logger.NewLogger()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	pgDSN := config.MakePGConn(cfg)
	conn, err := pgxpool.Connect(context.Background(), pgDSN)
	if err != nil {
		log.Fatalf("failed to open connection to posgtres: %s", err.Error())
	}

	roomRepository := roomRepo.NewRoomRepository(conn)

	userClientConn, err := grpc.Dial("chat-user:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to open connetion oto user service: %s", err)
	}

	userClient := desc.NewUserClient(userClientConn)

	roomService := room.NewRoomService(roomRepository, logger)
	chatService := chat.NewChatService(roomRepository, logger)

	handler := v1.NewHandler(roomService, userClient, chatService, logger)

	router := handler.Init()

	log.Printf("Starting router at port: %d...", cfg.HTTPPort)

	go chatService.Run()
	if err := router.Run(fmt.Sprintf(":%d", cfg.HTTPPort)); err != nil {
		log.Fatalf("failed to run router on port %s %s", cfg.HTTPPort, err.Error())
	}

}
