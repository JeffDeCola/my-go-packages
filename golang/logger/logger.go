package mylogger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

// The logLevel is passed as a string but I convert to int
// so I can upgrade the logger i use and it's just easier
type MyLogLevel string
type myLogLevelInt int

// My Levels as string
const (
	Trace   MyLogLevel = "trace"
	Debug   MyLogLevel = "debug"
	Info    MyLogLevel = "info"
	Warning MyLogLevel = "warning"
	Error   MyLogLevel = "error"
	Fatal   MyLogLevel = "fatal"
)

// My Levels as int
const (
	trace   myLogLevelInt = 0
	debug   myLogLevelInt = 1
	info    myLogLevelInt = 2
	warning myLogLevelInt = 3
	error   myLogLevelInt = 4
	fatal   myLogLevelInt = 5
)

// Map my log Levels to Slog Levels
var sLogLevels = map[myLogLevelInt]slog.Leveler{
	trace:   slog.LevelDebug,
	debug:   slog.LevelDebug,
	info:    slog.LevelInfo,
	warning: slog.LevelWarn,
	error:   slog.LevelError,
	fatal:   slog.LevelError,
}

// Formatting for jeffs format
var logLevelNames = map[myLogLevelInt]string{
	trace:   "TRACE",
	debug:   "DEBUG",
	info:    "INFO ",
	warning: "WARN ",
	error:   "ERROR",
	fatal:   "FATAL",
}

// Colors with for jeffs format
var logLevelColors = map[myLogLevelInt]string{
	trace:   "grey",
	debug:   "cyan",
	info:    "green",
	warning: "yellow",
	error:   "red",
	fatal:   "magenta",
}

// My logger struct
type theLoggerStruct struct {
	theSetLevel myLogLevelInt
	theFormat   string   // jeffs, jeffs_noTime, text, json
	theOutput   *os.File // stdout, stderr, filename
	theLogger   *slog.Logger
}

// CreateLogger
func CreateLogger(myLevel MyLogLevel, format string, output *os.File) *theLoggerStruct {

	// Convert string MyLogLevel to myLogLevelInt
	// Just easier dealing with  a number for the arrays
	myLevelInt := ConvertStringToLogLevel(myLevel)

	// If myLevel is not 0,1,2,3,4,5, then default to 2 (info)
	if myLevelInt < trace || myLevelInt > fatal {
		myLevelInt = info
	}

	// Create a handler with a log level
	var handler slog.Handler
	switch format {
	case "text":
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{
			Level: sLogLevels[myLevelInt]})
	case "json":
		handler = slog.NewJSONHandler(output, &slog.HandlerOptions{
			Level: sLogLevels[myLevelInt]})
	default:
		// Won't use handler, but i create it to put in struct
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{
			Level: sLogLevels[myLevelInt],
		})
	}

	// Create the logger struct or are we changing the myLevelInt
	l := &theLoggerStruct{
		theFormat:   format,
		theOutput:   output,
		theSetLevel: myLevelInt,
		theLogger:   slog.New(handler),
	}
	return l
}

// ChangeLogLevel changes the log level
func (l *theLoggerStruct) ChangeLogLevel(myLevel MyLogLevel) {

	myLevelInt := ConvertStringToLogLevel(myLevel)

	// If myLevel is not 0,1,2,3,4,5, then default to 2 (info)
	if myLevelInt < trace || myLevelInt > fatal {
		myLevelInt = info
	}

	// Update the log level of the existing logger
	l.theSetLevel = myLevelInt

	// Create a new handler with the updated log level
	var handler slog.Handler
	switch l.theFormat {
	case "text":
		handler = slog.NewTextHandler(l.theOutput, &slog.HandlerOptions{
			Level: sLogLevels[myLevelInt],
		})
	case "json":
		handler = slog.NewJSONHandler(l.theOutput, &slog.HandlerOptions{
			Level: sLogLevels[myLevelInt],
		})
	default:
		// Won't use handler, but create to put in struct
		handler = slog.NewTextHandler(l.theOutput, &slog.HandlerOptions{
			Level: sLogLevels[myLevelInt],
		})
	}

	// Update the logger with the new handler
	l.theLogger = slog.New(handler)

}

// ConvertStringToLogLevel converts a string log level to its corresponding myLogLevelInt value.
// It returns Info level as default for unrecognized values.
func ConvertStringToLogLevel(levelStr MyLogLevel) myLogLevelInt {
	var levelInt myLogLevelInt

	switch levelStr {
	case "Trace":
		levelInt = trace
	case "Debug":
		levelInt = debug
	case "Info":
		levelInt = info
	case "Warning":
		levelInt = warning
	case "Error":
		levelInt = error
	case "Fatal":
		levelInt = fatal
	default:
		levelInt = info // Default to Info if invalid
	}

	return levelInt
}

func (l *theLoggerStruct) Trace(msg string, args ...interface{}) {
	l.logMessage(trace, msg, args...)
}

func (l *theLoggerStruct) Debug(msg string, args ...interface{}) {
	l.logMessage(debug, msg, args...)
}

func (l *theLoggerStruct) Info(msg string, args ...interface{}) {
	l.logMessage(info, msg, args...)
}

func (l *theLoggerStruct) Warning(msg string, args ...interface{}) {
	l.logMessage(warning, msg, args...)
}

func (l *theLoggerStruct) Error(msg string, args ...interface{}) {
	l.logMessage(error, msg, args...)
}

func (l *theLoggerStruct) Fatal(msg string, args ...interface{}) {
	l.logMessage(Fatal, msg, args...)
	os.Exit(1)
}

// Print and format the message if needed
func (l *theLoggerStruct) logMessage(level myLogLevelInt, msg string, args ...any) {

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
func (l *theLoggerStruct) jeffsLogMessage(level myLogLevelInt, msg string, args ...any) {

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
