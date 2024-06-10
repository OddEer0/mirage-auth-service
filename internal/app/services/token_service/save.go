package tokenService

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
)

func (s *service) Save(ctx context.Context, data appDto.SaveTokenServiceDto) (*model.JwtToken, error) {
	stackTrace.Add(ctx, "package: tokenService, type: service, method: Save")
	defer stackTrace.Done(ctx)

	save, err := s.tokenRepository.Save(ctx, data.Id, data.RefreshToken)
	if err != nil {
		return nil, err
	}
	return save, nil
}
