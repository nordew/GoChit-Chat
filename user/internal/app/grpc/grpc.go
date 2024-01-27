package grpcApp

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"log/slog"
	"net"
	"user/internal/api/user"
	"user/internal/service"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(port int, gRPCServer *grpc.Server, userService service.UserService, logger *slog.Logger, zapLogger *zap.Logger) *App {
	// Don't log paylods in production =)
	//loggingOpts := []logging.Option{
	//	logging.WithLogOnEvents(
	//		logging.PayloadSent, logging.PayloadReceived,
	//	),
	//}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Fatalf("Recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	gRPCServer = grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		//logging.UnaryServerInterceptor(InterceptorLogger(logger), loggingOpts...)),
	))

	user.Register(gRPCServer, userService, zapLogger)

	return &App{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

//func InterceptorLogger(l *slog.Logger) logging.Logger {
//	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
//		l.Log(ctx, slog.Level(lvl), msg, fields...)
//	})
//}

func (a *App) Run() {
	const op = "grpcApp.App.Run"

	log.Printf("%s: starting gRPC server on port %d", op, a.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("%s: failed to listen: %v", op, err)
	}

	if err := a.gRPCServer.Serve(lis); err != nil {
		log.Fatalf("%s: failed to serve: %v", op, err)
	}
}

func (a *App) Stop() {
	const op = "grpcApp.App.Stop"

	log.Printf("%s: stopping gRPC server", op)

	a.gRPCServer.GracefulStop()

	log.Printf("%s: gRPC server stopped", op)
}
