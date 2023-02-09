# My circle package

_A package containing computations related to prime numbers._

Documentation and reference,

* Circle package at [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle)
* Refer to my example
  [module-with-remote-package](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/module-with-remote-package)

## TYPE & METHODS

Types,

```go
type Circle struct {
    R float64
}
```

Methods,

```go
func (c Circle) **circleArea**() float64
func (c Circle) **circleCircumference**() float64
```
