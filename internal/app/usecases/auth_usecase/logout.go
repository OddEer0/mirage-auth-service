package authUseCase

import (
	"context"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
)

func (u *useCase) Logout(ctx context.Context, refreshToken string) error {
	stackTrace.Add(ctx, "package: authUseCase, type: useCase, method: Logout")
	defer stackTrace.Done(ctx)

	err := u.tokenService.DeleteByValue(ctx, refreshToken)
	if err != nil {
		return err
	}
	return nil
}
