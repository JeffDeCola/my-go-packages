# My circle package

_A package to calculate the area and circumference of a circle._

Documentation and reference,

* Circle package at [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/geometry/circle)
* Refer to my example
  [module-with-remote-package](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/module-with-remote-package)

## TYPES

The Circle type is,

```go
type Circle struct {
    R float64
}
```

## METHODS

```go
func (c Circle) circleArea() float64
func (c Circle) circleCircumference() float64
```

## EXAMPLE

```go
package main

import (
    "fmt"
    "github.com/JeffDeCola/my-go-packages/geometry/circle"
)

func main() {

    // Create a circle
    c := circle.Circle{R: 5}

    // Calculate the area
    a := c.CircleArea()
    fmt.Println("Area =", a)

    // Calculate the circumference
    p := c.CircleCircumference()
    fmt.Println("Circumference =", p)
}
```
