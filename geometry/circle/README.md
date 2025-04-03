# MY CIRCLE PACKAGE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A package to calculate the area and circumference of a circle._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle#overview)
  * [TYPES](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle#types)
  * [METHODS](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle#methods)
* [EXAMPLE](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle#example)
* [ADD TO YOUR GO.MOD](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle#add-to-your-gomod)

Documentation and Reference

* Circle package at
  [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/geometry/circle)
* Refer to my example
  [module-with-remote-package](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/module-with-remote-package)

## OVERVIEW

This package contains a Circle type with two methods to calculate the area
and circumference of a circle.

### TYPES

```go
type Circle struct {
    Radius float64
}
```

### METHODS

```go
func (c Circle) Area() float64
func (c Circle) Circumference() float64
```

## EXAMPLE

```go
package main

import (
    "fmt"

    "github.com/JeffDeCola/my-go-packages/geometry/circle"
)

func main() {

    // Create a Circle type
    c := circle.Circle{Radius: 5}

    // Get the area
    a := c.Area()
    fmt.Println("Area =", a)

    // Get the circumference
    p := c.Circumference()
    fmt.Println("Circumference =", p)

}
```

## ADD TO YOUR GO.MOD

Since each package is tagged independently,

```text
git tag geometry/circle/vX.X.X
git push --tags
```

Add this to your go.mod file,

```text
require github.com/JeffDeCola/my-go-packages/geometry/circle vX.X.X
```
