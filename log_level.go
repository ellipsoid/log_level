// Package level_log provides leveled logging.
package log_level

type Level int

const (
	TRACE Level = iota
	DEBUG Level = iota
	INFO  Level = iota
	WARN  Level = iota
	ERROR Level = iota
)

// A Logger represents a logging object that will filter logging
// messages based on its internal log level and print filtered
// results to an object that meets the printer interface.
type Logger struct {
	level  Level
	logger printer
}

type printer interface {
	Printf(string, ...interface{})
}

// New creates a new Logger. The level variable sets the minimum
// threshold for logs to be output. The logger variable sets the
// destination to which log data will be written.
func New(level Level, logger printer) Logger {
	return Logger{level: level, logger: logger}
}

// Trace prints the formated string to the logger if the log level
// is TRACE or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Trace(format string, v ...interface{}) {
	if l.level <= TRACE {
		l.logger.Printf(format, v...)
	}
	return
}

// Debug prints the formated string to the logger if the log level
// is DEBUG or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Debug(format string, v ...interface{}) {
	if l.level <= DEBUG {
		l.logger.Printf(format, v...)
	}
	return
}

// Info prints the formated string to the logger if the log level
// is INFO or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Info(format string, v ...interface{}) {
	if l.level <= INFO {
		l.logger.Printf(format, v...)
	}
	return
}

// Warn prints the formated string to the logger if the log level
// is WARN or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Warn(format string, v ...interface{}) {
	if l.level <= WARN {
		l.logger.Printf(format, v...)
	}
	return
}

// Error prints the formated string to the logger if the log level
// is ERROR or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Error(format string, v ...interface{}) {
	if l.level <= ERROR {
		l.logger.Printf(format, v...)
	}
	return
}
