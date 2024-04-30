package grpcv1

import (
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
)

type AuthServiceServer struct {
	authv1.UnimplementedAuthServiceServer
}

func New() authv1.AuthServiceServer {
	return &AuthServiceServer{}
}
