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
  * A slog wrapper using NewTextHandler or JSON
  * 5 log levels to dynamically change
  * Custom formatting
  * Custom Color
* FUTURE ADDITIONS TO CONSIDER
  * Log to file

## CONST

```go
const (
    LevelDebug   myLogLevel = 0
    LevelInfo    myLogLevel = 1
    LevelWarning myLogLevel = 2
    LevelError   myLogLevel = 3
    LevelFatal   myLogLevel = 4
)
```

## TYPES

```go
type theLoggerStruct struct {
    theMode     string // jeffs, text, json
    theSetLevel myLogLevel
    theLogger   *slog.Logger
}
```

## FUNCTIONS

```go
func CreateLogger(myLevel myLogLevel, mode string) *theLoggerStruct {
```

## METHODS

```go
func (l *theLoggerStruct) ChangeLogLevel(myLevel myLogLevel) {
func (l *theLoggerStruct) Debug(msg string, args ...interface{}) {
func (l *theLoggerStruct) Info(msg string, v ...interface{}) {
func (l *theLoggerStruct) Warning(msg string, v ...interface{}) {
func (l *theLoggerStruct) Error(msg string, v ...interface{}) {
func (l *theLoggerStruct) Fatal(msg string, v ...interface{}) {

```

## EXAMPLE

```go
package main

import (
    "fmt"
    logger "my-go-packages/golang/logger"
)

func main() {
    log := logger.CreateLogger(logger.Debug, "jeffs")

    log.Debug("This is a debug message")
    log.Info("This is a Info Message", "env", "production", "user", "jeff")
    log.Warning("This is a Warning Message", "user", "jeff")
    log.Error("This is an Error message")
    // log.Fatal("Fatal Error")

    // Dynamically change log level
    fmt.Printf("\nCHANGE LEVEL\n\n")
    log.ChangeLogLevel(logger.Warning)

    log.Debug("This is a debug message")
    log.Info("This is a Info Message", "env", "production", "user", "jeff")
    log.Warning("This is a Warning Message", "user", "jeff")
    log.Error("This is an Error message")

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
