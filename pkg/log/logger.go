package log

import (
	"fmt"
	"log/slog"
	"os"
)

const (
	LevelTrace     = slog.Level(-8)
	LevelDebug     = slog.LevelDebug
	LevelInfo      = slog.LevelInfo
	LevelNotice    = slog.Level(2)
	LevelWarning   = slog.LevelWarn
	LevelError     = slog.LevelError
	LevelPanic     = slog.Level(9)
	LevelEmergency = slog.Level(12)
	LevelFatal     = slog.Level(16)
)

// CustomLogLevel replaces slog level names with custom names.
func CustomLogLevel(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		level := a.Value.Any().(slog.Level)
		switch {
		case level <= LevelTrace:
			a.Value = slog.StringValue("TRACE")
		case level <= LevelDebug:
			a.Value = slog.StringValue("DEBUG")
		case level <= LevelInfo:
			a.Value = slog.StringValue("INFO")
		case level <= LevelNotice:
			a.Value = slog.StringValue("NOTICE")
		case level <= LevelWarning:
			a.Value = slog.StringValue("WARNING")
		case level <= LevelError:
			a.Value = slog.StringValue("ERROR")
		case level <= LevelPanic:
			a.Value = slog.StringValue("PANIC")
		case level <= LevelEmergency:
			a.Value = slog.StringValue("EMERGENCY")
		default:
			a.Value = slog.StringValue("FATAL")
		}
	}
	return a
}

// Logger wraps slog with convenience methods.
type Logger struct {
	*slog.Logger
}

// NewLogger creates a new Logger with JSON output.
func NewLogger() *Logger {
	opts := &slog.HandlerOptions{
		Level:       LevelDebug,
		ReplaceAttr: CustomLogLevel,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	return &Logger{Logger: slog.New(handler)}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Log(nil, LevelDebug, fmt.Sprintf(format, args...))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Log(nil, LevelInfo, fmt.Sprintf(format, args...))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Log(nil, LevelWarning, fmt.Sprintf(format, args...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Log(nil, LevelError, fmt.Sprintf(format, args...))
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Log(nil, LevelFatal, fmt.Sprintf(format, args...))
	os.Exit(1)
}

func (l *Logger) Panic(msg string) {
	l.Log(nil, LevelPanic, msg)
	panic(msg)
}
