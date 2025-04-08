# MY SHAPES PACKAGE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A package to calculate and manipulate simple 2D and 3D geometric shapes._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/shapes#overview)
  * [INTERFACES](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/shapes#interfaces)
  * [TYPES](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/shapes#types)
  * [METHODS](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/shapes#methods)
  * [FUNCTIONS](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/shapes#functions)
* [EXAMPLE](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/shapes#example)
* [ADD TO YOUR GO.MOD](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/shapes#add-to-your-gomod)

Documentation and Reference

* shapes package at
  [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/geometry/shapes)
* refer to my example
  [shapes-package](https://github.com/JeffDeCola/my-go-examples/tree/master/functions-methods-interfaces/interfaces/shapes-package)

## OVERVIEW

This package contains a few types and interfaces to calculate the area, perimeter,
volume, and surface area of simple 2D and 3D geometric shapes. It also has a function
to change the size of the shapes.

### INTERFACES

```go
type TwoDCalculations interface {
    Area() float64
    Perimeter() float64
}
type ThreeDCalculations interface {
    Volume() float64
    SurfaceArea() float64
}
type ShapeManipulations interface {
    Size(f float64)
}
```

### TYPES

```go
type Rectangle struct {
    Width  float64
    Height float64
}
type Circle struct {
    Radius float64
}
type Triangle struct {
    Base   float64
    Height float64
}
type Cube struct {
    Length float64
}
type Sphere struct {
    Radius float64
}

```

### METHODS

```go
func (r Rectangle) Area() float64
func (r Rectangle) Perimeter() float64
func (c Circle) Area() float64
func (c Circle) Perimeter() float64
func (t Triangle) Area() float64
func (t Triangle) Perimeter() float64
func (c Cube) Volume() float64
func (c Cube) SurfaceArea() float64
func (s Sphere) Volume() float64
func (s Sphere) SurfaceArea() float64
```

```go
func (r *Rectangle) Size(f float64)
func (c *Circle) Size(f float64)
func (t *Triangle) Size(f float64)
func (c *Cube) Size(f float64)
func (s *Sphere) Size(f float64)
```

### FUNCTIONS

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

## ADD TO YOUR GO.MOD

Since each package is tagged independently,

```text
git tag geometry/shapes/vX.X.X
git push --tags
```

Add this to your go.mod file,

```text
require github.com/JeffDeCola/my-go-packages/geometry/shapes vX.X.X
```
