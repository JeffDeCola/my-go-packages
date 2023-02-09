# My shapes package

_A package to calculate and manipulate simple 2D and 3D geometric shapes._

Documentation and reference,

* Shapes package at [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/geometry/shapes)
* Refer to my example
  [geometry-package](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/geometry-package)

## INTERFACE

I'll just post the interface for a rectangle,

```go
TYPE

Types,

```go
type Geometry interface {
    Area(*float64)
    Perimeter(*float64)
    Size(float64)
}
```

For example, a Rectangle type would be,

```go
type Rectangle struct {
    Width  float64
    Height float64
}
```

Where it's Methods are,

```go
func (r Rectangle) Area(a *float64)
func (r Rectangle) Perimeter(p *float64)
func (r *Rectangle) Size(f float64)
```

```go
package main

import (
    "fmt"
    "github.com/JeffDeCola/my-go-packages/geometry/shapes"
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
