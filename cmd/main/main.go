package main

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/logger"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres"
	grpcv1AuthService "github.com/OddEer0/mirage-auth-service/internal/presentation/handler/grpcv1/grpcv1_auth_service"
	"github.com/OddEer0/mirage-auth-service/internal/presentation/interactor"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

func main() {
	cfg := config.MustLoad()
	log := logger.MustLoad(cfg.Env, cfg.Path.LogFile)
	conn, err := postgres.Connect(cfg, log)
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Error("close postgres error", slog.String("cause", err.Error()))
		}
	}(conn, context.Background())
	if err != nil {
		log.Error("sql connect error", slog.String("cause", err.Error()))
		return
	}
	lis, err := net.Listen("tcp", cfg.Server.Address)
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			log.Error("close tcp connection error", slog.String("cause", err.Error()))
		}
	}(lis)
	if err != nil {
		log.Error("net listen tcp error", slog.String("cause", err.Error()))
		return
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(logger.LoggingInterceptor),
	)

	dependencies := interactor.New(cfg, log, conn)

	authv1.RegisterAuthServiceServer(grpcServer, grpcv1AuthService.New(&grpcv1AuthService.Dependencies{
		UserRepository: dependencies.UserRepository,
		AuthUseCase:    dependencies.AuthUseCase,
		Log:            log,
	}))
	if err := grpcServer.Serve(lis); err != nil {
		log.Error("grpc serve error", "cause", err.Error())
	}
}
