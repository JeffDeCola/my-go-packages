# MY LOGGER PACKAGE

_A package that uses the standard library
[slog](https://pkg.go.dev/log/slog)
and enhances it for my liking._

Table of Contents

* [FUNCTIONS](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#functions)
* [EXAMPLE](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger#example)

Documentation and Reference

* Prime package at
  [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/golang/logger)

## FUNCTIONS

```go
func SetupLogger(level slog.Level) *slog.Logger {
```

## EXAMPLE

```go
package main

import (
    "fmt"

    "github.com/JeffDeCola/my-go-packages/golang/logger"
)

func main() {

    log := logger.SetupLogger(slog.LevelInfo) // slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError

    log.Debug("This is a debug message")
    log.Info("Application started")
    log.Warn("This is a warning")
    log.Error("An error occurred")

}
```
