package logger

import (
	"io"
	logDefault "log"
	"log/slog"
	"os"
)

const (
	EnvLocal   = "local"
	EnvDev     = "dev"
	EnvProd    = "prod"
	EnvTest    = "test"
	stdoutPath = "stdout"
)

var log *slog.Logger

func MustLoad(env string, filePath string) *slog.Logger {
	if log != nil {
		return log
	}

	out := os.Stdout
	if filePath != "" && filePath != stdoutPath {
		file, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
		if err != nil {
			logDefault.Fatalf("log file does not exist: %s", err)
		}
		out = file
	}

	switch env {
	case EnvLocal:
		log = setupLocalLog(out)
	case EnvDev:
		log = setupDevLog(out)
	case EnvProd:
		log = setupProdLog(out)
	case EnvTest:
		log = setupTestLog(out)
	}

	return log
}

func setupLocalLog(out io.Writer) *slog.Logger {
	jsonHandler := slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelDebug})
	return slog.New(jsonHandler)
}

func setupDevLog(out io.Writer) *slog.Logger {
	handler := AppLogHandler{slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelInfo})}
	return slog.New(handler)
}

func setupProdLog(out io.Writer) *slog.Logger {
	jsonHandler := slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelError})
	return slog.New(jsonHandler)
}

func setupTestLog(out io.Writer) *slog.Logger {
	handler := AppLogHandler{slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelError})}
	return slog.New(handler)
}
