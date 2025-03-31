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
* Example using this logger in
  [my-go-examples](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/jeffs-logger)

## OVERVIEW

You have three choices when you create a logger,

* myLevel (int)
  * Trace
  * Debug
  * Info **(default)**
  * Warning
  * Error
  * Fatal
* format (string)
  * "text" - Uses slog text
  * "json" - Uses slog json
  * "jeffs" - My own formatting **(default)**
  * "jeffs_noTime" - My own formatting without time stamp
* output (*os.File)
  * os.Stdout
  * os.Stderr
  * file handler

For example,

```go
    log := logger.CreateLogger(logger.Debug, "json", os.Stdout)
```

### CONST

```go
const (
    LevelTrace   myLogLevel = 0
    LevelDebug   myLogLevel = 1
    LevelInfo    myLogLevel = 2
    LevelWarning myLogLevel = 3
    LevelError   myLogLevel = 4
    LevelFatal   myLogLevel = 5
)
```

### TYPES

```go
type theLoggerStruct struct {
    theSetLevel myLogLevel
    theFormat   string   // jeffs, jeffs_noTime, text, json
    theOutput   *os.File // stdout, stderr, filename
    theLogger   *slog.Logger
}
```

### FUNCTIONS

```go
func CreateLogger(myLevel myLogLevel, format string, output *os.File) *theLoggerStruct {
```

### METHODS

```go
func (l *theLoggerStruct) ChangeLogLevel(myLevel myLogLevel) {
func (l *theLoggerStruct) Trace(msg string, args ...interface{}) {
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

    log := logger.CreateLogger(logger.Debug, "json", os.Stdout)

    a := 4.54534

    log.Trace("This is a low level trace message")
    log.Debug("This is a debug message")
    log.Info(fmt.Sprintf("Formatted Info Message a=%.2f", a), "a", a, "user", "jeff")
    log.Warning("This is a Warning Message", "user", "jeff")
    log.Error("This is an Error message")
    // log.Fatal("Fatal Error")

    // Dynamically change log level
    fmt.Printf("\nCHANGE LEVEL\n\n")
    log.ChangeLogLevel(logger.Warning)

    log.Trace("This is a low level trace message")
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
