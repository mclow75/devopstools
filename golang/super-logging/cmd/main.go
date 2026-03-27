package main

import (
	_ "cmd/main.go/internal/logging"
	"errors"
	"log/slog"
	"os"
)

func main() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(_.Logger)

	Logger.Info("Starting the application")
	Logger.Error("Error occurred", slog.Any("error", errors.New("test error")))
}
