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
                l.printMessage("[TRACE] ", format)
	}
	return
}

// Debug prints the formated string to the logger if the log level
// is DEBUG or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Debug(format string, v ...interface{}) {
	if l.level <= DEBUG {
                l.printMessage("[DEBUG] ", format)
	}
	return
}

// Info prints the formated string to the logger if the log level
// is INFO or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Info(format string, v ...interface{}) {
	if l.level <= INFO {
                l.printMessage("[INFO] ", format)
	}
	return
}

// Warn prints the formated string to the logger if the log level
// is WARN or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Warn(format string, v ...interface{}) {
	if l.level <= WARN {
                l.printMessage("[WARN] ", format)
	}
	return
}

// Error prints the formated string to the logger if the log level
// is ERROR or lower.
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Error(format string, v ...interface{}) {
	if l.level <= ERROR {
                l.printMessage("[ERROR] ", format)
	}
	return
}

func (l Logger) printMessage(prefix string, format string, v ...interface{}) {
        prefixFormat := prefix + format + "\n"
        l.logger.Printf(prefixFormat, v...)
	return
}
