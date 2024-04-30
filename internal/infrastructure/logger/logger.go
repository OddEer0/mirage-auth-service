package logger

import (
	"io"
	"log/slog"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
	EnvTest  = "test"
)

var log *slog.Logger

func SetupLogger(env string) *slog.Logger {
	if log != nil {
		return log
	}

	switch env {
	case EnvLocal:
		log = setupLocalLog(os.Stdout)
	case EnvDev:
		log = setupDevLog(os.Stdout)
	case EnvProd:
		log = setupProdLog(os.Stdout)
	case EnvTest:
		log = setupTestLog(os.Stdout)
	}

	return log
}

func setupLocalLog(out io.Writer) *slog.Logger {
	jsonHandler := slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelDebug})
	return slog.New(jsonHandler)
}

func setupDevLog(out io.Writer) *slog.Logger {
	jsonHandler := slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelInfo})
	return slog.New(jsonHandler)
}

func setupProdLog(out io.Writer) *slog.Logger {
	jsonHandler := slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelError})
	return slog.New(jsonHandler)
}

func setupTestLog(out io.Writer) *slog.Logger {
	jsonHandler := slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelError})
	return slog.New(jsonHandler)
}
