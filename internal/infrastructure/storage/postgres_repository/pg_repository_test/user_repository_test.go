package pg_repository_test

import (
	postgresRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository"
	mockgenPostgres "github.com/OddEer0/mirage-auth-service/internal/tests/mockgen/pg_mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserPgRepository(t *testing.T) {
	//mockUserData := mock2.PostgresData().User
	//ctx := context.Background()
	//mock := getMockPg(ctx, mockUserData)
	//defer func() {
	//	_ = mock.Close(ctx)
	//}()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mockgenPostgres.NewMockQuery(ctrl)
	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByIdQuery)
	//tLog := logger.MustLoad("test", "")
	//userRepo := postgresRepository.NewUserRepository(tLog, mock)
	//
	//t.Run("Should correct get by id", func(t *testing.T) {
	//	user := mockUserData.CorrectUser1
	//	userDb, err := userRepo.GetById(ctx, user.Id)
	//	require.NoError(t, err)
	//	assert.Equal(t, user, userDb)
	//
	//	user = mockUserData.BannedUser1
	//	userDb, err = userRepo.GetById(ctx, user.Id)
	//	require.NoError(t, err)
	//	assert.Equal(t, user, userDb)
	//
	//	user = mockUserData.AdminUser1
	//	userDb, err = userRepo.GetById(ctx, user.Id)
	//	require.NoError(t, err)
	//	assert.Equal(t, user, userDb)
	//})

	//t.Run("Should correct create user", func(t *testing.T) {
	//	createUser := mockUserData.CreateUser1Res
	//	userDb, err := userRepo.Create(ctx, createUser)
	//	require.NoError(t, err)
	//	assert.Equal(t, createUser, userDb)
	//})
}
