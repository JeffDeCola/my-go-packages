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
	Trace   myLogLevel = 0
	Debug   myLogLevel = 1
	Info    myLogLevel = 2
	Warning myLogLevel = 3
	Error   myLogLevel = 4
	Fatal   myLogLevel = 5
)

// Map my log Levels to Slog Levels
var sLogLevels = map[myLogLevel]slog.Leveler{
	Trace:   slog.LevelDebug,
	Debug:   slog.LevelDebug,
	Info:    slog.LevelInfo,
	Warning: slog.LevelWarn,
	Error:   slog.LevelError,
	Fatal:   slog.LevelError,
}

// Formatting for jeffs format
var logLevelNames = map[myLogLevel]string{
	Trace:   "TRACE",
	Debug:   "DEBUG",
	Info:    "INFO ",
	Warning: "WARN ",
	Error:   "ERROR",
	Fatal:   "FATAL",
}

// Colors with for jeffs format
var logLevelColors = map[myLogLevel]string{
	Trace:   "grey",
	Debug:   "cyan",
	Info:    "green",
	Warning: "yellow",
	Error:   "red",
	Fatal:   "magenta",
}

// My logger struct
type theLoggerStruct struct {
	theSetLevel myLogLevel
	theFormat   string   // jeffs, jeffs_noTime, text, json
	theOutput   *os.File // stdout, stderr, filename
	theLogger   *slog.Logger
}

// CreateLogger
func CreateLogger(myLevel myLogLevel, format string, output *os.File) *theLoggerStruct {

	// If myLevel is not 0,1,2,3,4,5, then default to 2 (info)
	if myLevel < Trace || myLevel > Fatal {
		myLevel = Info
	}

	// Create a handler with a log level
	var handler slog.Handler
	switch format {
	case "text":
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{
			Level: sLogLevels[myLevel]})
	case "json":
		handler = slog.NewJSONHandler(output, &slog.HandlerOptions{
			Level: sLogLevels[myLevel]})
	default:
		// Won't use handler, but i create it to put in struct
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{
			Level: sLogLevels[myLevel],
		})
	}

	// Create the logger struct or are we changing the myLevel
	l := &theLoggerStruct{
		theFormat:   format,
		theOutput:   output,
		theSetLevel: myLevel,
		theLogger:   slog.New(handler),
	}
	return l
}

// ChangeLogLevel changes the log level
func (l *theLoggerStruct) ChangeLogLevel(myLevel myLogLevel) {

	// If myLevel is not 0,1,2,3,4,5, then default to 2 (info)
	if myLevel < Trace || myLevel > Fatal {
		myLevel = Info
	}

	// Update the log level of the existing logger
	l.theSetLevel = myLevel

	// Create a new handler with the updated log level
	var handler slog.Handler
	switch l.theFormat {
	case "text":
		handler = slog.NewTextHandler(l.theOutput, &slog.HandlerOptions{
			Level: sLogLevels[myLevel],
		})
	case "json":
		handler = slog.NewJSONHandler(l.theOutput, &slog.HandlerOptions{
			Level: sLogLevels[myLevel],
		})
	default:
		// Won't use handler, but create to put in struct
		handler = slog.NewTextHandler(l.theOutput, &slog.HandlerOptions{
			Level: sLogLevels[myLevel],
		})
	}

	// Update the logger with the new handler
	l.theLogger = slog.New(handler)

}

func (l *theLoggerStruct) Trace(msg string, args ...interface{}) {
	l.logMessage(Trace, msg, args...)
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

	switch l.theFormat {
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

	// Get the color for level name
	color := logLevelColors[level]
	colorCode := getColorCode(color)
	levelName = fmt.Sprintf("\033[%sm%s\033[0m", colorCode, levelName)

	// get the time
	theTime := time.Now().Format("15:04:05")

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
	// Print the message to l.theOutput (stderr or stdout)
	if l.theFormat == "jeffs_noTime" {
		fmt.Fprintf(l.theOutput, "[%s] %s %s\n", levelName, message, theArgs)
	} else {
		fmt.Fprintf(l.theOutput, "[%s] [%s] %s %s\n", theTime, levelName, message, theArgs)
	}

}

func getColorCode(color string) string {
	switch color {
	case "grey":
		return "90" // Grey
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
