package internal

import (
	"log/slog"
	"os"
)

func SetupLogger() {
	logLevelVar := loadEnvVar("LOG_LEVEL", "DEBUG")
	var logLevel slog.Leveler
	if logLevelVar == "DEBUG" {
		logLevel = slog.LevelDebug
	} else {
		logLevel = slog.LevelInfo
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	slog.SetDefault(logger)
}
