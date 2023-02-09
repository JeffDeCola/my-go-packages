# My shapes package

_A package to calculate and manipulate simple 2D and 3D geometric shapes._

Documentation and reference,

* Shapes package at [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/geometry/shapes)
* Refer to my example
  [geometry-package](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/geometry-package)

## FUNCTIONS

I'll just post the functions a user can use,

```go
func GetArea(t TwoDCalculations) float64
func GetPerimeter(t TwoDCalculations) float64
func GetVolume(t ThreeDCalculations) float64
func GetSurfaceArea(t ThreeDCalculations) float64
func ChangeSize(t ShapeManipulations, f float64)
```

## EXAMPLE

```go
package main

import (
    "fmt"
    "github.com/JeffDeCola/my-go-packages/geometry/shapoes"
)

func main() {

    // Create a circle type
    c := circle.Circle{Radius: 5}

    // Get the area
    a := c.Area()
    fmt.Println("Area =", a)

    // Get the circumference
    p := c.Circumference()
    fmt.Println("Circumference =", p)
}
```

Where go.mod is,

```text
module module-with-remote-package

go 1.19

require github.com/JeffDeCola/my-go-packages v0.0.9
```
