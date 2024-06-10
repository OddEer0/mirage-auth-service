package userUseCase

import (
	"context"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
)

func (u *useCase) DeleteById(ctx context.Context, id string) error {
	stackTrace.Add(ctx, "package: userUseCase, type: useCase, method: DeleteById")
	defer stackTrace.Done(ctx)

	return u.userRepository.Delete(ctx, id)
}
