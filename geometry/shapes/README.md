# My shapes package

_A package to calculate and manipulate simple 2D and 3D geometric shapes._

Documentation and reference,

* Shapes package at
  [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/geometry/shapes)
* Refer to my example
  [shapes-package](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/shapes-package)

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

    "github.com/JeffDeCola/my-go-packages/geometry/shapes"
)

func main() {

    // Create a Rectangle and Circle type
    r := shapes.Rectangle{Width: 10, Height: 10}
    c := shapes.Circle{Radius: 5}

    // Get the area (using the interface)
    a := shapes.GetArea(r)
    fmt.Println("Area of rectangle =", a)
    a = shapes.GetArea(c)
    fmt.Println("Area of circle =", a)

    // Change the size (x2)
    shapes.ChangeSize(&r, 2)
    shapes.ChangeSize(&c, 2)

    // Get the area (using the interface)
    a = shapes.GetArea(r)
    fmt.Println("Area of rectangle =", a)
    a = shapes.GetArea(c)
    fmt.Println("Area of circle =", a)

}
```

Where go.mod is,

```text
module shapes-package

go 1.19

require github.com/JeffDeCola/my-go-packages v0.2.0
```
