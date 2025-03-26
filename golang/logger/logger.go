package mylogger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

// The logLevel is just an integer so I can upgrade loggers later
type myLogLevel int

// My Levels
const (
	LevelDebug   myLogLevel = 0
	LevelInfo    myLogLevel = 1
	LevelWarning myLogLevel = 2
	LevelError   myLogLevel = 3
	LevelFatal   myLogLevel = 4
)

// Map my log Levels to Slog Levels
var sLogLevels = map[myLogLevel]slog.Leveler{
	LevelDebug:   slog.LevelDebug,
	LevelInfo:    slog.LevelInfo,
	LevelWarning: slog.LevelWarn,
	LevelError:   slog.LevelError,
	LevelFatal:   slog.LevelError,
}

// Formatting
var logLevelNames = map[myLogLevel]string{
	LevelDebug:   "DEBUG",
	LevelInfo:    "INFO ",
	LevelWarning: "WARN ",
	LevelError:   "ERROR",
	LevelFatal:   "FATAL",
}

// Colors
var logLevelColors = map[myLogLevel]string{
	LevelDebug:   "cyan",
	LevelInfo:    "green",
	LevelWarning: "yellow",
	LevelError:   "red",
	LevelFatal:   "magenta",
}

// My logger struct
type theLoggerStruct struct {
	theSetLevel myLogLevel // Don't really need this, but why not
	theLogger   *slog.Logger
}

func CreateLogger(level myLogLevel) *theLoggerStruct {

	// Get the slog handler struct
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: sLogLevels[level]})

	// Create the struct to pass to main
	l := &theLoggerStruct{
		theSetLevel: level,
		theLogger:   slog.New(handler),
	}
	return l
}

func (l *theLoggerStruct) ChangeLogLevel(level myLogLevel) {
	l.theSetLevel = level

	// Must get a new handler
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: sLogLevels[level]})
	l.theLogger = slog.New(handler)
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

func (l *theLoggerStruct) logMessage(level myLogLevel, msg string, args ...any) {

	// Map requested LogLevel to slog.Level
	slogLevel := sLogLevels[level].Level()

	// Add the current time as a formatted string
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Format the log level name
	levelName := logLevelNames[level]

	// Add color formatting (optional, for terminal output)
	//color := logLevelColors[level]
	//formattedLevel := fmt.Sprintf("\033[%sm%s\033[0m", colorCode(color), levelName)

	// Format the log message as a single string
	formattedMessage := fmt.Sprintf("[%s] %s: %s", currentTime, levelName, msg)

	// Append additional arguments if provided
	if len(args) > 0 {
		formattedMessage = fmt.Sprintf("%s | %v", formattedMessage, args)
	}

	// Output the formatted message using slog
	l.theLogger.Log(context.Background(), slogLevel, formattedMessage)

}

// Helper function to map color names to ANSI color codes
func colorCode(color string) string {
	switch color {
	case "cyan":
		return "36"
	case "green":
		return "32"
	case "yellow":
		return "33"
	case "red":
		return "31"
	case "magenta":
		return "35"
	default:
		return "0" // Default color
	}
}
