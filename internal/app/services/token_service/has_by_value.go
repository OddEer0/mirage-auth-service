package tokenService

import (
	"context"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
)

func (s *service) HasByValue(ctx context.Context, refreshToken string) (bool, error) {
	stackTrace.Add(ctx, "package: tokenService, type: service, method: Save")
	defer stackTrace.Done(ctx)

	has, err := s.tokenRepository.HasByValue(ctx, refreshToken)
	if err != nil {
		return false, err
	}

	return has, nil
}
