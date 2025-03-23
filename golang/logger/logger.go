package mylogger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

// The logLevel is just an integer so I can upgrade loggers later
type LogLevel int

// My Levels
const (
	LevelDebug   LogLevel = 0
	LevelInfo    LogLevel = 1
	LevelWarning LogLevel = 2
	LevelError   LogLevel = 3
	LevelFatal   LogLevel = 4
)

// Slog Mapping
var sLogLevels = map[LogLevel]slog.Leveler{
	LevelDebug:   slog.LevelDebug,
	LevelInfo:    slog.LevelInfo,
	LevelWarning: slog.LevelWarn,
	LevelError:   slog.LevelError,
	LevelFatal:   slog.LevelError,
}

// For Formatting
var logLevelNames = map[LogLevel]string{
	LevelDebug:   "DEBUG",
	LevelInfo:    "INFO ",
	LevelWarning: "WARN ",
	LevelError:   "ERROR",
	LevelFatal:   "FATAL",
}

// Colors
var logLevelColors = map[LogLevel]string{
	LevelDebug:   "cyan",
	LevelInfo:    "green",
	LevelWarning: "yellow",
	LevelError:   "red",
	LevelFatal:   "magenta",
}

// My logger struct
type theLoggerStruct struct {
	theSetLevel LogLevel // Don't really need this, but why not
	slogLogger  *slog.Logger
}

func CreateLogger(level LogLevel) *theLoggerStruct {

	var handler slog.Handler

	// Get the slog handler struct
	handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: sLogLevels[level]})

	// Create the struct to pass to main
	l := &theLoggerStruct{
		theSetLevel: level,
		slogLogger:  slog.New(handler),
	}
	return l
}

func (l *theLoggerStruct) ChangeLogLevel(level LogLevel) {
	l.theSetLevel = level

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: sLogLevels[level]})
	l.slogLogger = slog.New(handler)
}

func (l *theLoggerStruct) Debug(message string, v ...interface{}) {
	l.logMessage(LevelDebug, message, v...)
}

func (l *theLoggerStruct) Info(message string, v ...interface{}) {
	l.logMessage(LevelInfo, message, v...)
}

func (l *theLoggerStruct) Warning(message string, v ...interface{}) {
	l.logMessage(LevelWarning, message, v...)
}

func (l *theLoggerStruct) Error(message string, v ...interface{}) {
	l.logMessage(LevelError, message, v...)
}

func (l *theLoggerStruct) Fatal(message string, v ...interface{}) {
	l.logMessage(LevelFatal, message, v...)
}

// log handles the actual slog output
func (l *theLoggerStruct) logMessage(level LogLevel, msg string, args ...any) {

	// Map requested LogLevel to slog.Level
	var slogLevel slog.Level
	switch level {
	case LevelDebug:
		fmt.Println("D")
		slogLevel = slog.LevelDebug
	case LevelInfo:
		fmt.Println("I")
		slogLevel = slog.LevelInfo
	case LevelWarning:
		fmt.Println("W")
		slogLevel = slog.LevelWarn
	case LevelError, LevelFatal:
		fmt.Println("E")
		slogLevel = slog.LevelError
	}

	// Add the current time as a structured attribute
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Log the message using slog with structured arguments
	l.slogLogger.Log(context.Background(), slogLevel, msg, "time", currentTime, "args", args)

}
