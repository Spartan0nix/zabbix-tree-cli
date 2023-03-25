package logging

import (
	"fmt"
	"log"
)

// LogLevel is used to defined the cricity level of the logger.
type LogLevel int

const (
	Critical LogLevel = iota
	Error
	Warning
	Info
	Debug
)

// Logger is used as a wrapper to handle custom shell logging format.
// Logging to a file is not supported.
type Logger struct {
	Level  LogLevel
	logger *log.Logger
}

// NewLogger is used to create a new logger.
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		Level:  level,
		logger: log.Default(),
	}
}

// getLevel is used to returned the string representation of a LogLevel value.
func getLevel(level LogLevel) string {
	switch level {
	case Critical:
		return "CRITICAL"
	case Error:
		return "ERROR"
	case Warning:
		return "WARNING"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	default:
		return "UNSUPPORTED"
	}
}

// setFlags is used to set the logger flags based on the given log level.
func (l *Logger) setFlags(level LogLevel) {
	if level <= Error {
		l.logger.SetFlags(log.Lshortfile | log.Lmsgprefix)
	} else {
		l.logger.SetFlags(log.Lmsgprefix)
	}
}

// log is used as a wrapper function to log the given data while handling logger flags, prefix and content assignments.
func (l *Logger) log(level LogLevel, v ...any) {
	// Do not write log if the level is not configured
	if level <= l.Level {
		// Set the logger flags
		l.setFlags(level)

		// Retrieve the string representation of the log level
		strLevel := getLevel(level)

		// For critical level, log all data to the shell and
		// execute panic function on the last entry
		if level == Critical {
			// Configure the logger prefix
			l.logger.SetPrefix(fmt.Sprintf("[%s] ", strLevel))

			// Directly use panic function if only one entry was passed
			if len(v) == 1 {
				panic(v[0])
			}

			// Log all entries to the shell, minus the latest one
			for i := 0; i < len(v)-1; i++ {
				l.logger.Println(v[i])
			}
			panic(v[len(v)-1])

		} else {
			// Configure the logger prefix
			l.logger.SetPrefix(fmt.Sprintf("[%s] ", strLevel))
			// Log all entries to the shell
			for _, entry := range v {
				l.logger.Println(entry)
			}
		}
	}
}

// Critical is used to log data to the shell with the level logging.Critical.
func (l *Logger) Critical(v ...any) {
	l.log(Critical, v...)
}

// Error is used to log data to the shell with the level logging.Error.
func (l *Logger) Error(v ...any) {
	l.log(Error, v...)
}

// Warning is used to log data to the shell with the level logging.Warning.
func (l *Logger) Warning(v ...any) {
	l.log(Warning, v...)
}

// Info is used to log data to the shell with the level logging.Info.
func (l *Logger) Info(v ...any) {
	l.log(Info, v...)
}

// Debug is used to log data to the shell with the level logging.Debug.
func (l *Logger) Debug(v ...any) {
	l.log(Debug, v...)
}
