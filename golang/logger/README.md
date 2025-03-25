# MY LOGGER PACKAGE

_Just a logger wrapper I use for my liking.
Currently using the standard library
[slog](https://pkg.go.dev/log/slog)
 with NewTextHandler (Not JSON)._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#overview)
* [FUNCTIONS](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#functions)
* [EXAMPLE](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#example)

Documentation and Reference

* Prime package at
  [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/golang/logger)

## OVERVIEW

* CURRENT FEATURES
  * A slog wrapper using NewTextHandler
  * 5 Log Levels
  * Custom formatting
  * Custom Color
* FUTURE ADDITIONS TO CONSIDER
  * JSON support using New JSON Handler
  * Key value logging
  * Log to file

## METHODS

```go
func CreateLogger(level LogLevel) *theLoggerStruct {
func (l *theLoggerStruct) ChangeLogLevel(level LogLevel) {
func (l *theLoggerStruct) Debug(message string, v ...interface{}) {
func (l *theLoggerStruct) Info(message string, v ...interface{}) {
func (l *theLoggerStruct) Warning(message string, v ...interface{}) {
func (l *theLoggerStruct) Error(message string, v ...interface{}) {
func (l *theLoggerStruct) Fatal(message string, v ...interface{}) {

```

## EXAMPLE

```go
package main

import (
    "fmt"

    "github.com/JeffDeCola/my-go-packages/golang/logger"
)

func main() {

    log := logger.CreateLogger(logger.LevelDebug)
    // slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError

    log.Debug("This is a debug message")
    log.Info("Application started")
    log.Warn("This is a warning")
    log.Error("An error occurred")
    log.Fatal("Fatal, not good")

    log.ChangeLogLevel(logger.LevelWarning)

    log.Debug("This is a debug message") // Won't Print
    log.Info("Application started")      // Won't print
    log.Warn("This is a warning")
    log.Error("An error occurred")
    log.Fatal("Fatal, not good")

}
```

## ADDED TO YOUR GO.MOD

Since I am tagging each package independently,

```text
git tag golang/logger/vX.X.X
git push --tags
```

This will be added to your go.mod file,

```text
require github.com/JeffDeCola/my-go-packages/golang/logger vX.X.X
```
