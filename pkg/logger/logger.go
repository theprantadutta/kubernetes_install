package logger

import (
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

const logTimeFormat = "2006-01-02 15:04:05"

// Logger defines a simple logger with different log levels
type Logger struct {
	log *log.Logger
}

// New creates a new instance of Logger
func New() *Logger {
	return &Logger{
		log: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info logs an informational message with color and timestamp
func (l *Logger) Info(msg string, args ...interface{}) {
	infoLog(msg, args...)
}

// Success logs a success message with color and timestamp
func (l *Logger) Success(msg string, args ...interface{}) {
	successLog(msg, args...)
}

// Error logs an error message with color and timestamp
func (l *Logger) Error(msg string, args ...interface{}) {
	errorLog(msg, args...)
}

// Helper functions for logging with colors and timestamps
func infoLog(message string, args ...interface{}) {
	currentTime := time.Now().Format(logTimeFormat)
	color.HiCyan("[INF] "+currentTime+" "+message, args...)
}

func successLog(message string, args ...interface{}) {
	currentTime := time.Now().Format(logTimeFormat)
	color.HiGreen("[SUC] "+currentTime+" "+message, args...)
}

func errorLog(message string, args ...interface{}) {
	currentTime := time.Now().Format(logTimeFormat)
	color.HiRed("[ERR] "+currentTime+" "+message, args...)
}
