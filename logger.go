// Package log_level provides leveled logging.
package log_level

// A Logger represents a logging object that will filter
// messages based on its internal log level. It prints filtered
// results to its logger object, which must satisfy the printer interface.
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
		prefixFormat := "[TRACE] " + format
		l.logger.Printf(prefixFormat, v...)
	}
	return
}

// Debug prints the formated string to the logger if the log level
// is DEBUG or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Debug(format string, v ...interface{}) {
	if l.level <= DEBUG {
		prefixFormat := "[DEBUG] " + format
		l.logger.Printf(prefixFormat, v...)
	}
	return
}

// Info prints the formated string to the logger if the log level
// is INFO or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Info(format string, v ...interface{}) {
	if l.level <= INFO {
		prefixFormat := "[INFO] " + format
		l.logger.Printf(prefixFormat, v...)
	}
	return
}

// Warn prints the formated string to the logger if the log level
// is WARN or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Warn(format string, v ...interface{}) {
	if l.level <= WARN {
		prefixFormat := "[WARN] " + format
		l.logger.Printf(prefixFormat, v...)
	}
	return
}

// Error prints the formated string to the logger if the log level
// is ERROR or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Error(format string, v ...interface{}) {
	if l.level <= ERROR {
		prefixFormat := "[ERROR] " + format
		l.logger.Printf(prefixFormat, v...)
	}
	return
}
