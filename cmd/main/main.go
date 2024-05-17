package main

import (
	"context"
	"fmt"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/logger"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres"
	"github.com/OddEer0/mirage-auth-service/internal/presentation/handler/grpcv1"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"net"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
	conn, err := postgres.Connect(cfg, log)
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			return
		}
	}(conn, context.Background())
	if err != nil {
		fmt.Println("sql connect error: ", err.Error())
		return
	}
	lis, err := net.Listen("tcp", cfg.Server.Address)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(logger.LoggingInterceptor),
	)
	authv1.RegisterAuthServiceServer(grpcServer, grpcv1.New(cfg, log, conn))
	log.Info("Server started to address: " + cfg.Server.Address)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
