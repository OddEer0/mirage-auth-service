package mock

import (
	domainConstants "github.com/OddEer0/mirage-auth-service/internal/domain/domain_constants"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	"time"
)

type (
	Users struct {
		AdminUser1     *model.User
		CorrectUser1   *model.User
		BannedUser1    *model.User
		CreateUser1Res *model.User
		NotFoundUser   *model.User
		InternalUser   *model.User
		InternalUser2  *model.User
	}

	Tokens struct {
		AdminUser1   *model.JwtToken
		CorrectUser1 *model.JwtToken
		BannedUser1  *model.JwtToken
	}

	UserActivate struct {
		AdminUser1   *model.UserActivate
		CorrectUser1 *model.UserActivate
		BannedUser1  *model.UserActivate
	}

	Postgres struct {
		User         *Users
		Token        *Tokens
		UserActivate *UserActivate
	}
)

func PostgresData() *Postgres {
	banReason1 := "toxic"
	return &Postgres{
		User: &Users{
			CorrectUser1: &model.User{
				Id:        "111",
				Login:     "aboba",
				Email:     "bibas@gmail.com",
				Password:  "SuperSecretPass123",
				Role:      domainConstants.RoleUser,
				IsBanned:  false,
				BanReason: nil,
				UpdatedAt: time.Now().AddDate(-1, 0, 0),
				CreatedAt: time.Now().AddDate(-3, 0, 0),
			},
			BannedUser1: &model.User{
				Id:        "112",
				Login:     "toxic",
				Email:     "toxus@gmail.com",
				Password:  "SuperToxicPass123",
				Role:      domainConstants.RoleUser,
				IsBanned:  true,
				BanReason: &banReason1,
				UpdatedAt: time.Now().AddDate(-1, 0, 0),
				CreatedAt: time.Now().AddDate(-2, 0, 0),
			},
			AdminUser1: &model.User{
				Id:        "113",
				Login:     "Marlen",
				Email:     "merlin@gmail.com",
				Password:  "SuperAdminPass123",
				Role:      domainConstants.RoleAdmin,
				IsBanned:  false,
				BanReason: nil,
				UpdatedAt: time.Now().AddDate(-1, 0, 0),
				CreatedAt: time.Now().AddDate(-2, 0, 0),
			},
			CreateUser1Res: &model.User{
				Id:        "114",
				Login:     "General",
				Email:     "qingyuan@gmail.com",
				Password:  "SuperGeneralPass123",
				Role:      domainConstants.RoleUser,
				IsBanned:  false,
				BanReason: nil,
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			},
			NotFoundUser: &model.User{
				Id:    "not-found",
				Login: "not-found",
				Email: "not-found",
			},
			InternalUser: &model.User{
				Id:    "internal",
				Login: "internal",
				Email: "internal",
			},
			InternalUser2: &model.User{
				Id:    "internal2",
				Login: "internal2",
				Email: "internal2",
			},
		},
	}
}
