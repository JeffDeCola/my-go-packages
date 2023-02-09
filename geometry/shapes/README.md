# My shapes package

_A package to calculate and manipulate simple 2D and 3D geometric shapes._

Documentation and reference,

* Shapes package at [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/geometry/shapes)
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
?????????????????
```

Where go.mod is,

```text
module shapes-package

go 1.19

require github.com/JeffDeCola/my-go-packages v0.1.0
```
