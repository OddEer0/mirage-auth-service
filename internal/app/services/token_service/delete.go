package tokenService

import (
	"context"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
)

func (s *service) DeleteByValue(ctx context.Context, value string) error {
	stackTrace.Add(ctx, "package: tokenService, type: service, method: DeleteByValue")
	defer stackTrace.Done(ctx)

	err := s.tokenRepository.DeleteByValue(ctx, value)
	if err != nil {
		return err
	}

	return nil
}
