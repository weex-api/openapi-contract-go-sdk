package weex

import (
	"fmt"
	"log"
	"os"
)

// LogLevel represents the logging level
type LogLevel int

const (
	LogLevelDebug LogLevel = iota // Debug level - most verbose
	LogLevelInfo                  // Info level - general information
	LogLevelWarn                  // Warn level - warning messages
	LogLevelError                 // Error level - error messages
	LogLevelNone                  // None - no logging
)

// String returns the string representation of LogLevel
func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	case LogLevelNone:
		return "NONE"
	default:
		return "UNKNOWN"
	}
}

// Logger is the interface for logging in the SDK
type Logger interface {
	// Debug logs a debug message
	Debug(msg string, args ...interface{})

	// Info logs an info message
	Info(msg string, args ...interface{})

	// Warn logs a warning message
	Warn(msg string, args ...interface{})

	// Error logs an error message
	Error(msg string, args ...interface{})

	// SetLevel sets the logging level
	SetLevel(level LogLevel)
}

// DefaultLogger is the default logger implementation using Go's standard log package
type DefaultLogger struct {
	level  LogLevel
	logger *log.Logger
}

// NewDefaultLogger creates a new default logger with the specified log level
func NewDefaultLogger(level LogLevel) *DefaultLogger {
	return &DefaultLogger{
		level:  level,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Debug logs a debug message
func (l *DefaultLogger) Debug(msg string, args ...interface{}) {
	if l.level <= LogLevelDebug {
		l.log("DEBUG", msg, args...)
	}
}

// Info logs an info message
func (l *DefaultLogger) Info(msg string, args ...interface{}) {
	if l.level <= LogLevelInfo {
		l.log("INFO", msg, args...)
	}
}

// Warn logs a warning message
func (l *DefaultLogger) Warn(msg string, args ...interface{}) {
	if l.level <= LogLevelWarn {
		l.log("WARN", msg, args...)
	}
}

// Error logs an error message
func (l *DefaultLogger) Error(msg string, args ...interface{}) {
	if l.level <= LogLevelError {
		l.log("ERROR", msg, args...)
	}
}

// SetLevel sets the logging level
func (l *DefaultLogger) SetLevel(level LogLevel) {
	l.level = level
}

// log is the internal logging method
func (l *DefaultLogger) log(level, msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	l.logger.Printf("[%s] %s", level, msg)
}

// NoOpLogger is a logger that does nothing
type NoOpLogger struct{}

// NewNoOpLogger creates a new no-op logger
func NewNoOpLogger() *NoOpLogger {
	return &NoOpLogger{}
}

// Debug does nothing
func (l *NoOpLogger) Debug(msg string, args ...interface{}) {}

// Info does nothing
func (l *NoOpLogger) Info(msg string, args ...interface{}) {}

// Warn does nothing
func (l *NoOpLogger) Warn(msg string, args ...interface{}) {}

// Error does nothing
func (l *NoOpLogger) Error(msg string, args ...interface{}) {}

// SetLevel does nothing
func (l *NoOpLogger) SetLevel(level LogLevel) {}
