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
	Debug   myLogLevel = 0
	Info    myLogLevel = 1
	Warning myLogLevel = 2
	Error   myLogLevel = 3
	Fatal   myLogLevel = 4
)

// Map my log Levels to Slog Levels
var sLogLevels = map[myLogLevel]slog.Leveler{
	Debug:   slog.LevelDebug,
	Info:    slog.LevelInfo,
	Warning: slog.LevelWarn,
	Error:   slog.LevelError,
	Fatal:   slog.LevelError,
}

// Formatting with jeff mode
var logLevelNames = map[myLogLevel]string{
	Debug:   "DEBUG",
	Info:    "INFO ",
	Warning: "WARN ",
	Error:   "ERROR",
	Fatal:   "FATAL",
}

// Colors with jeff mode
var logLevelColors = map[myLogLevel]string{
	Debug:   "cyan",
	Info:    "green",
	Warning: "yellow",
	Error:   "red",
	Fatal:   "magenta",
}

// My logger struct
type theLoggerStruct struct {
	theMode     string // jeffs, text, json
	theSetLevel myLogLevel
	theLogger   *slog.Logger
}

// CreateTextLogger
func CreateLogger(myLevel myLogLevel, mode string) *theLoggerStruct {

	// Create a handler with a log level
	var handler slog.Handler
	switch mode {
	case "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: sLogLevels[myLevel]})
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: sLogLevels[myLevel]})
	default:
		// Won't use handler anyway
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: sLogLevels[myLevel],
		})
	}

	// Create the logger struct or are we changing the myLevel
	l := &theLoggerStruct{
		theMode:     mode,
		theSetLevel: myLevel,
		theLogger:   slog.New(handler),
	}
	return l
}

// Just update struct with never info
func (l *theLoggerStruct) ChangeLogLevel(myLevel myLogLevel) {

	// Update the log level of the existing logger
	l.theSetLevel = myLevel

	// Create a new handler with the updated log level
	var handler slog.Handler
	switch l.theMode {
	case "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: sLogLevels[myLevel],
		})
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: sLogLevels[myLevel],
		})
	default:
		// Won't use handler anyway
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: sLogLevels[myLevel],
		})
	}

	// Update the logger with the new handler
	l.theLogger = slog.New(handler)

}

func (l *theLoggerStruct) Debug(msg string, args ...interface{}) {
	l.logMessage(Debug, msg, args...)
}

func (l *theLoggerStruct) Info(msg string, args ...interface{}) {
	l.logMessage(Info, msg, args...)
}

func (l *theLoggerStruct) Warning(msg string, args ...interface{}) {
	l.logMessage(Warning, msg, args...)
}

func (l *theLoggerStruct) Error(msg string, args ...interface{}) {
	l.logMessage(Error, msg, args...)
}

func (l *theLoggerStruct) Fatal(msg string, args ...interface{}) {
	l.logMessage(Fatal, msg, args...)
	os.Exit(1)
}

// Print and format the message if needed
func (l *theLoggerStruct) logMessage(level myLogLevel, msg string, args ...any) {

	switch l.theMode {
	case "text":
		// Could add other info in the msg
		l.theLogger.Log(context.Background(), sLogLevels[level].Level(), msg, args...)

	case "json":
		// Could add other info in the msg
		l.theLogger.Log(context.Background(), sLogLevels[level].Level(), msg, args...)

	default:
		// default to jeffs
		l.jeffsLogMessage(level, msg, args...)
	}

}

// jeffs Log Message
func (l *theLoggerStruct) jeffsLogMessage(level myLogLevel, msg string, args ...any) {

	// only print the current level
	if level < l.theSetLevel {
		return
	}

	// Get the level name
	levelName := logLevelNames[level]

	// get the time
	time := time.Now().Format("15:04:05")

	// Get the color
	color := logLevelColors[level]
	colorCode := getColorCode(color)

	// get message
	message := msg

	// Get the args (key value pairs)
	var theArgs string
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) { // Ensure there is a value for the key
			theArgs += fmt.Sprintf("(%v: %v)", args[i], args[i+1])
			if i+2 < len(args) { // Add a comma only if there are more pairs
				theArgs += ", "
			}
		}
	}
	// print the message
	// fmt.Printf("%s %s %s %s\n", levelName, time, message, theArgs)
	fmt.Printf("[\033[%sm%s\033[0m] %s %s %s\n", colorCode, levelName, time, message, theArgs)

}

func getColorCode(color string) string {
	switch color {
	case "cyan":
		return "36" // Cyan
	case "green":
		return "32" // Green
	case "yellow":
		return "33" // Yellow
	case "red":
		return "31" // Red
	case "magenta":
		return "35" // Magenta
	default:
		return "0" // Default (no color)
	}
}
