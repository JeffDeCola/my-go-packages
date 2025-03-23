package mylogger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

type LogLevel int

const (
	LevelDebug   LogLevel = 0
	LevelInfo    LogLevel = 1
	LevelWarning LogLevel = 2
	LevelError   LogLevel = 3
	LevelFatal   LogLevel = 4
)

var sLogLevels = map[LogLevel]slog.Leveler{
	LevelDebug:   slog.LevelDebug,
	LevelInfo:    slog.LevelInfo,
	LevelWarning: slog.LevelWarn,
	LevelError:   slog.LevelError,
	LevelFatal:   slog.LevelError,
}

var logLevelNames = map[LogLevel]string{
	LevelDebug:   "DEBUG",
	LevelInfo:    "INFO ",
	LevelWarning: "WARN ",
	LevelError:   "ERROR",
	LevelFatal:   "FATAL",
}

var logLevelColors = map[LogLevel]string{
	LevelDebug:   "cyan",
	LevelInfo:    "green",
	LevelWarning: "yellow",
	LevelError:   "red",
	LevelFatal:   "magenta",
}

type theLoggerStruct struct {
	theSetLevel LogLevel // Don't really need this
	slogLogger  *slog.Logger
}

func CreateLogger(level LogLevel) *theLoggerStruct {

	var handler slog.Handler

	// Text output with custom formatting
	handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: sLogLevels[level],
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case "time":
				a.Value = slog.StringValue(time.Now().Format("2006-01-02 15:04:05"))
			case "level":
				levelMap := map[slog.Level]string{
					slog.LevelDebug: "DEBUG",
					slog.LevelInfo:  "INFO",
					slog.LevelWarn:  "WARN",
					slog.LevelError: "ERROR",
				}
				a.Value = slog.StringValue(levelMap[a.Value.Any().(slog.Level)])
			}
			return a
		},
	})

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
