# MY LOGGER PACKAGE

_A package that uses the standard library
[slog](https://pkg.go.dev/log/slog) NewTextHandler (Not JSON)
and enhances it for my liking._

Table of Contents

* [FUNCTIONS](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#functions)
* [EXAMPLE](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#example)

Documentation and Reference

* Prime package at
  [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/golang/logger)

## FUNCTIONS

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

}
```
