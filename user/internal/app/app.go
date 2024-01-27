package app

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	grpcApp "user/internal/app/grpc"
	"user/internal/config"
	userRepo "user/internal/repository/user"
	"user/internal/service/user"
	bcryptHasher "user/pkg/hasher/bcrypt"
)

func MustRun(ctx context.Context) {
	app := injectDependencies(ctx)
	app.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	app.Stop()
}

func injectDependencies(ctx context.Context) *grpcApp.App {
	const op = "app.injectDependencies"

	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to create zap logger: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		zapLogger.Error("failed to load .env file", zap.Error(err), zap.String("op", op))
	}

	cfg, err := config.GetConfig()
	if err != nil {
		zapLogger.Error("failed to get config", zap.Error(err), zap.String("op", op))
	}

	pool, err := setupDatabase(ctx, cfg)
	if err != nil {
		zapLogger.Error("failed to setup database", zap.Error(err), zap.String("op", op))
	}

	userRepository := userRepo.NewUserRepository(pool, zapLogger)
	hasher := bcryptHasher.NewPasswordHasher()
	userService := user.NewUserService(userRepository, hasher, zapLogger)

	grpcServer := grpc.NewServer()

	logger := setupLogger()

	app := grpcApp.New(cfg.GRPCPort, grpcServer, userService, logger, zapLogger)

	return app
}

func setupDatabase(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	connStr := config.MakePGConn(cfg)

	pool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func setupLogger() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
}
