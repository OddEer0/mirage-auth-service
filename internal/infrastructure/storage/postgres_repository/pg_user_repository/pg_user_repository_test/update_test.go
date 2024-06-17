package pg_user_repository_test

import (
	"database/sql"
	"errors"
	pgUserRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository/pg_user_repository"
	"github.com/OddEer0/mirage-auth-service/internal/tests/mock"
	mockgenPostgres "github.com/OddEer0/mirage-auth-service/internal/tests/mockgen/mockgen_postgres"
	testCtx "github.com/OddEer0/mirage-auth-service/internal/tests/test_ctx"
	testLogger "github.com/OddEer0/mirage-auth-service/internal/tests/test_logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUpdate(t *testing.T) {
	mockPgData := mock.PostgresData()
	mockUserData := mockPgData.User

	t.Run("Testing UpdateById", func(t *testing.T) {
		tLog := testLogger.New()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		mockDb.EXPECT().QueryRow(
			gomock.Any(),
			pgUserRepository.UpdateUserById,
			mockUserData.CorrectUser1.Id,
			"new_login",
			mockUserData.CorrectUser1.Email,
			mockUserData.CorrectUser1.Password,
			mockUserData.CorrectUser1.Role,
			mockUserData.CorrectUser1.IsBanned,
			mockUserData.CorrectUser1.BanReason,
			mockUserData.CorrectUser1.UpdatedAt,
		).Return(mock.PgMockRow{
			Data: []any{
				mockUserData.CorrectUser1.Id,
				"new_login",
				mockUserData.CorrectUser1.Email,
				mockUserData.CorrectUser1.Password,
				mockUserData.CorrectUser1.Role,
				mockUserData.CorrectUser1.IsBanned,
				mockUserData.CorrectUser1.BanReason,
				mockUserData.CorrectUser1.UpdatedAt,
				mockUserData.CorrectUser1.CreatedAt,
			},
		})

		mockDb.EXPECT().QueryRow(
			gomock.Any(),
			pgUserRepository.UpdateUserById,
			mockUserData.NotFoundUser.Id,
			mockUserData.NotFoundUser.Login,
			mockUserData.NotFoundUser.Email,
			mockUserData.NotFoundUser.Password,
			mockUserData.NotFoundUser.Role,
			mockUserData.NotFoundUser.IsBanned,
			mockUserData.NotFoundUser.BanReason,
			mockUserData.NotFoundUser.UpdatedAt,
		).Return(mock.PgMockRowError{Err: sql.ErrNoRows})

		mockDb.EXPECT().QueryRow(
			gomock.Any(),
			pgUserRepository.UpdateUserById,
			mockUserData.InternalUser.Id,
			mockUserData.InternalUser.Login,
			mockUserData.InternalUser.Email,
			mockUserData.InternalUser.Password,
			mockUserData.InternalUser.Role,
			mockUserData.InternalUser.IsBanned,
			mockUserData.InternalUser.BanReason,
			mockUserData.InternalUser.UpdatedAt,
		).Return(mock.PgMockRowError{Err: errors.New("internal")})

		userRepo := pgUserRepository.New(tLog, mockDb)

		t.Run("Should correct work", func(t *testing.T) {
			ctx := testCtx.New()
			user := *mockUserData.CorrectUser1
			user.Login = "new_login"
			userDb, err := userRepo.UpdateById(ctx, &user)
			require.NoError(t, err)
			assert.Equal(t, userDb, &user)
		})

		t.Run("Should not found", func(t *testing.T) {
			ctx := testCtx.New()
			user := *mockUserData.NotFoundUser
			userDb, err := userRepo.UpdateById(ctx, &user)
			assert.Nil(t, userDb)
			assert.Equal(t, pgUserRepository.ErrUserNotFound, err)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{pgUserRepository.TraceUpdateById}, tLog.Stack)
			tLog.Clean()
		})

		t.Run("Should not found", func(t *testing.T) {
			ctx := testCtx.New()
			user := *mockUserData.InternalUser
			userDb, err := userRepo.UpdateById(ctx, &user)
			assert.Nil(t, userDb)
			assert.Equal(t, pgUserRepository.ErrInternal, err)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{pgUserRepository.TraceUpdateById}, tLog.Stack)
			tLog.Clean()
		})
	})

	t.Run("Testing UpdatePasswordById", func(t *testing.T) {
		tLog := testLogger.New()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		updatedPassword := "update41_pass"
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.UpdateUserPasswordById, mockUserData.CorrectUser1.Id, updatedPassword).Return(
			mock.PgMockRow{Data: []any{
				mockUserData.CorrectUser1.Id,
				mockUserData.CorrectUser1.Login,
				mockUserData.CorrectUser1.Email,
				updatedPassword,
				mockUserData.CorrectUser1.Role,
				mockUserData.CorrectUser1.IsBanned,
				mockUserData.CorrectUser1.BanReason,
				mockUserData.CorrectUser1.UpdatedAt,
				mockUserData.CorrectUser1.CreatedAt,
			}})

		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.UpdateUserPasswordById, mockUserData.NotFoundUser.Id, updatedPassword).Return(
			mock.PgMockRowError{Err: sql.ErrNoRows})

		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.UpdateUserPasswordById, mockUserData.InternalUser.Id, updatedPassword).Return(
			mock.PgMockRowError{Err: errors.New("internal")})

		userRepo := pgUserRepository.New(tLog, mockDb)

		t.Run("Should correct work", func(t *testing.T) {
			ctx := testCtx.New()
			user := *mockUserData.CorrectUser1
			user.Password = updatedPassword
			userDb, err := userRepo.UpdatePasswordById(ctx, user.Id, updatedPassword)
			require.NoError(t, err)
			assert.Equal(t, &user, userDb)
		})

		t.Run("Should not found", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.NotFoundUser
			userDb, err := userRepo.UpdatePasswordById(ctx, user.Id, updatedPassword)
			assert.Nil(t, userDb)
			assert.Equal(t, pgUserRepository.ErrUserNotFound, err)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{pgUserRepository.TraceUpdatePasswordById}, tLog.Stack)
			tLog.Clean()
		})

		t.Run("Should not found", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			userDb, err := userRepo.UpdatePasswordById(ctx, user.Id, updatedPassword)
			assert.Nil(t, userDb)
			assert.Equal(t, pgUserRepository.ErrInternal, err)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{pgUserRepository.TraceUpdatePasswordById}, tLog.Stack)
			tLog.Clean()
		})
	})

	t.Run("Testing UpdateRoleById", func(t *testing.T) {
		tLog := testLogger.New()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		updatedRole := "aboba"
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.UpdateUserRoleById, mockUserData.CorrectUser1.Id, updatedRole).Return(
			mock.PgMockRow{Data: []any{
				mockUserData.CorrectUser1.Id,
				mockUserData.CorrectUser1.Login,
				mockUserData.CorrectUser1.Email,
				mockUserData.CorrectUser1.Password,
				updatedRole,
				mockUserData.CorrectUser1.IsBanned,
				mockUserData.CorrectUser1.BanReason,
				mockUserData.CorrectUser1.UpdatedAt,
				mockUserData.CorrectUser1.CreatedAt,
			}})

		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.UpdateUserRoleById, mockUserData.NotFoundUser.Id, updatedRole).Return(
			mock.PgMockRowError{Err: sql.ErrNoRows})

		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.UpdateUserRoleById, mockUserData.InternalUser.Id, updatedRole).Return(
			mock.PgMockRowError{Err: errors.New("internal")})

		userRepo := pgUserRepository.New(tLog, mockDb)

		t.Run("Should correct work", func(t *testing.T) {
			ctx := testCtx.New()
			user := *mockUserData.CorrectUser1
			user.Role = updatedRole
			userDb, err := userRepo.UpdateRoleById(ctx, user.Id, updatedRole)
			require.NoError(t, err)
			assert.Equal(t, &user, userDb)
		})

		t.Run("Should not found", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.NotFoundUser
			userDb, err := userRepo.UpdateRoleById(ctx, user.Id, updatedRole)
			assert.Nil(t, userDb)
			assert.Equal(t, pgUserRepository.ErrUserNotFound, err)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{pgUserRepository.TraceUpdateRoleById}, tLog.Stack)
			tLog.Clean()
		})

		t.Run("Should not found", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			userDb, err := userRepo.UpdateRoleById(ctx, user.Id, updatedRole)
			assert.Nil(t, userDb)
			assert.Equal(t, pgUserRepository.ErrInternal, err)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{pgUserRepository.TraceUpdateRoleById}, tLog.Stack)
			tLog.Clean()
		})
	})
}
