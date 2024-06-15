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

	TraceGetById            = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "GetById"}
	TraceGetByLogin         = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "GetByLogin"}
	TraceCreate             = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "Create"}
	TraceCheckUserRole      = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "CheckUserRole"}
	TraceGetByQuery         = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "GetByQuery"}
	TraceDelete             = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "Delete"}
	TraceUpdateById         = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "UpdateById"}
	TraceUpdateRoleById     = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "UpdateRoleById"}
	TraceUpdatePasswordById = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "UpdatePasswordById"}
	TraceBanUserById        = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "BanUserById"}
	TraceUnbanUserById      = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "UnbanUserById"}
	TraceHasUserById        = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "HasUserById"}
	TraceHasUserByLogin     = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "HasUserByLogin"}
	TraceHasUserByEmail     = stacktrace.Method{Package: "postgresRepository", Type: "postgresRepository", Method: "HasUserByEmail"}
)

const (
	LogDbQueryMessage = "error database query"
	LogNoRowMessage   = "row not found"
)
