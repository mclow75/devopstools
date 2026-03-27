// internal/logging/logging.go
package logging

import "log/slog"

var Logger slog.Logger

// LogUserAction logs a user-specific action with structured attributes.
func LogUserAction(userID int, action string) {
	// Use the default slog logger configured in the main package
	slog.Info("User action performed",
		"user_id", userID,
		"action", action,
	)
}

// LogError logs an error with a specific context.
func LogError(err error, context string) {
	slog.Error("An error occurred in internal package",
		"error", err, // slog handles errors well
		"context", context,
	)
}

// LogDebug can be used for detailed debugging information.
// It will only be output if the main package configured the logger for Debug level.
func LogDebug(msg string, args ...any) {
	slog.Debug(msg, args...)
}
