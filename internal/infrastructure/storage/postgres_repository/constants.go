package postgresRepository

import (
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
)

var (
	ErrUserActivateNotFound = domain.NewErr(domain.ErrNotFoundCode, "user activate not found")
	ErrTokenNotFound        = domain.NewErr(domain.ErrNotFoundCode, "token not found")
	ErrUserNotFound         = domain.NewErr(domain.ErrNotFoundCode, "user not found")
	ErrInternal             = domain.NewErr(domain.ErrInternalCode, "internal error")

	TraceUserRepoGetById            = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "GetById"}
	TraceUserRepoGetByLogin         = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "GetByLogin"}
	TraceUserRepoCreate             = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "Create"}
	TraceUserRepoCheckUserRole      = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "CheckUserRole"}
	TraceUserRepoGetByQuery         = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "GetByQuery"}
	TraceUserRepoDelete             = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "Delete"}
	TraceUserRepoUpdateById         = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "UpdateById"}
	TraceUserRepoUpdateRoleById     = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "UpdateRoleById"}
	TraceUserRepoUpdatePasswordById = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "UpdatePasswordById"}
	TraceUserRepoBanUserById        = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "BanUserById"}
	TraceUserRepoUnbanUserById      = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "UnbanUserById"}
	TraceUserRepoHasUserById        = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "HasUserById"}
	TraceUserRepoHasUserByLogin     = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "HasUserByLogin"}
	TraceUserRepoHasUserByEmail     = stacktrace.Method{Package: "postgresRepository", Type: "userRepository", Method: "HasUserByEmail"}
)

const (
	LogDbQueryMessage = "error database query"
	LogNoRowMessage   = "row not found"
)
