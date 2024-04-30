package grpcv1

import (
	"context"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
)

func (a *AuthServiceServer) LinkActivate(ctx context.Context, link *authv1.Link) (*authv1.Empty, error) {
	//TODO implement me
	panic("implement me")
}
