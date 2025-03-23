# MY CIRCLE PACKAGE

_A package to calculate the area and circumference of a circle._

Table of Contents

* [TYPES](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle#types)
* [METHODS](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle#methods)
* [EXAMPLE](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle#example)

Documentation and Reference

* Circle package at
  [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/geometry/circle)
* Refer to my example
  [module-with-remote-package](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/module-with-remote-package)

## TYPES

```go
type Circle struct {
    Radius float64
}
```

## METHODS

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

Where your go.mod is,

```text
module yourModuleName

go 1.24

require github.com/JeffDeCola/my-go-packages/geometry/circle v0.0.0
```
