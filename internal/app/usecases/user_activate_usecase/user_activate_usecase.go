package userActivateUsecase

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
)

type (
	UseCase interface {
		ActivateLink(ctx context.Context, link string) error
		IsUserActivate(ctx context.Context, userId string) (bool, error)
	}

	useCase struct {
		userActivateRepository repository.UserActivateRepository
	}
)

func (u *useCase) ActivateLink(ctx context.Context, link string) error {
	stackTrace.Add(ctx, "package: userActivateUseCase, type: useCase, method: ActivateLink")
	defer stackTrace.Done(ctx)

	_, err := u.userActivateRepository.ActivateUserByLink(ctx, link)
	return err
}

func (u *useCase) IsUserActivate(ctx context.Context, userId string) (bool, error) {
	stackTrace.Add(ctx, "package: userActivateUseCase, type: useCase, method: IsUserActivate")
	defer stackTrace.Done(ctx)

	isActivate, err := u.userActivateRepository.IsActivateById(ctx, userId)

	if err != nil {
		return false, err
	}
	return isActivate, nil
}

func New(activateRepository repository.UserActivateRepository) UseCase {
	return &useCase{
		userActivateRepository: activateRepository,
	}
}
