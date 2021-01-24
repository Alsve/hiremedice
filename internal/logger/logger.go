package logger

import (
	"fmt"
)

// L is global instance of the logger
var L = &StdOutLogger{}

// LoggerStdOut logs to standard out
type StdOutLogger struct{}

// Debug logs message at DEBUG level
func (l StdOutLogger) Debug(msg string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+msg, args...)
}

// Info logs message at INFO level
func (l StdOutLogger) Info(msg string, args ...interface{}) {
	fmt.Printf("[INFO] "+msg, args...)
}

// Warn logs message at WARN level
func (l StdOutLogger) Warn(msg string, args ...interface{}) {
	fmt.Printf("[WARN] "+msg, args...)
}

// Error logs message at ERROR level
func (l StdOutLogger) Error(msg string, args ...interface{}) {
	fmt.Printf("[ERROR] "+msg, args...)
}
