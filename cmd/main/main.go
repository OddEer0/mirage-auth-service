package main

import (
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	"github.com/OddEer0/mirage-auth-service/internal/presentation/handler/grpcv1"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
	"google.golang.org/grpc"
	"net"
)

func main() {
	cfg := config.MustLoad()
	lis, err := net.Listen("tcp", cfg.Server.Address)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	authv1.RegisterAuthServiceServer(grpcServer, grpcv1.New())
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
