# MY LOGGER PACKAGE

_Just a logger wrapper I use for my liking.
Currently using the standard library
[slog](https://pkg.go.dev/log/slog)._

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
* LOG LEVELS
  * Debug
  * Info
  * Warning
  * Error
  * Fatal
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
