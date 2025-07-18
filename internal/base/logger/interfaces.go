package logger

// Logger is an interface for logging messages.
type Logger interface {
	// Info logs provided arguments at INFO level.
	Info(args ...any)
	// Infof formats the message according to the format specifier and logs it at INFO level.
	Infof(template string, args ...any)
	// Debug logs provided arguments at DEBUG level.
	Debug(args ...any)
	// Debugf formats the message according to the format specifier and logs it at DEBUG level.
	Debugf(template string, args ...any)
	// Warn logs provided arguments at WARN level.
	Warn(args ...any)
	// Warnf formats the message according to the format specifier and logs it at WARN level.
	Warnf(template string, args ...any)
	// Panic logs provided arguments at PANIC level.
	Panic(args ...any)
	// Panicf formats the message according to the format specifier and logs it at PANIC level.
	Panicf(template string, args ...any)
	// Error logs provided arguments at ERROR level.
	Error(args ...any)
	// Errorf formats the message according to the format specifier and logs it at ERROR level.
	Errorf(template string, args ...any)
}
