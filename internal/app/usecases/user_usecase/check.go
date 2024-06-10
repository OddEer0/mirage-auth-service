package userUseCase

import (
	"context"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
)

func (u *useCase) CheckUserRole(ctx context.Context, userId, role string) (bool, error) {
	stackTrace.Add(ctx, "package: userUseCase, type: useCase, method: CheckUserRole")
	defer stackTrace.Done(ctx)

	return u.userRepository.CheckUserRole(ctx, userId, role)
}
