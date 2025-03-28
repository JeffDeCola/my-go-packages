# MY LOGGER PACKAGE

_Just a logger wrapper formatting a message followed
by key-value pairs.
Currently using the standard library
[slog](https://pkg.go.dev/log/slog)
supporting both text and json._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#overview)
  * [CONST](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#const)
  * [TYPES](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#types)
  * [FUNCTIONS](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#functions)
  * [METHODS](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#methods)
* [EXAMPLE](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#example)
* [ADD TO YOUR GO.MOD](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#add-to-your-gomod)

Documentation and Reference

* Prime package at
  [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/golang/logger)

## OVERVIEW

* MODES OF OPERATION
  * "text" - Uses slog text
  * "json" - Uses slog json
  * "jeffs" - My own formatting
  * "jeffs_noTime" - My own formatting without time
* LOG LEVELS
  * Debug
  * Info
  * Warning
  * Error
  * Fatal
* THE OUTPUT
  * Time
  * Leg Level
  * Message
  * Key/Value Pairs
* FUTURE ADDITIONS TO CONSIDER
  * Log to file

### CONST

```go
const (
    LevelDebug   myLogLevel = 0
    LevelInfo    myLogLevel = 1
    LevelWarning myLogLevel = 2
    LevelError   myLogLevel = 3
    LevelFatal   myLogLevel = 4
)
```

### TYPES

```go
type theLoggerStruct struct {
    theMode     string // jeffs, text, json
    theSetLevel myLogLevel
    theLogger   *slog.Logger
}
```

### FUNCTIONS

```go
func CreateLogger(myLevel myLogLevel, mode string) *theLoggerStruct {
```

### METHODS

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
    logger "github.com/JeffDeCola/my-go-packages/golang/logger"
)

func main() {

    log := logger.CreateLogger(logger.Debug, "jeffs")

    a := 4.54534

    log.Debug("This is a debug message")
    log.Info(fmt.Sprintf("Formatted Info Message a=%.2f", a), "a", a, "user", "jeff")
    log.Warning("This is a Warning Message", "user", "jeff")
    log.Error("This is an Error message")
    // log.Fatal("Fatal Error")

    // Dynamically change log level
    fmt.Printf("\nCHANGE LEVEL\n\n")
    log.ChangeLogLevel(logger.Warning)

    log.Debug("This is a debug message")
    log.Info(fmt.Sprintf("Formatted Info Message a=%.2f", a), "a", a, "user", "jeff")
    log.Warning("This is a Warning Message", "user", "jeff")
    log.Error("This is an Error message")

}
```

## ADD TO YOUR GO.MOD

Since each package is tagged independently,

```text
git tag golang/logger/vX.X.X
git push --tags
```

Add this to your go.mod file,

```text
require github.com/JeffDeCola/my-go-packages/golang/logger vX.X.X
```
