package main

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
	"google.golang.org/grpc"
	"net"
)

type AuthServiceServer struct {
	authv1.UnimplementedAuthServiceServer
}

func (a *AuthServiceServer) Registration(ctx context.Context, request *authv1.RegistrationRequest) (*authv1.AuthResponse, error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	cfg := config.MustLoad()
	lis, err := net.Listen("tcp", cfg.Server.Address)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	authv1.RegisterAuthServiceServer(grpcServer, &AuthServiceServer{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
