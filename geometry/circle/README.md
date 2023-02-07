# My circle package

_My circle package contains computations for area and circumference of a circle._

Table of Contents,

* [TYPE & METHODS](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle#type--methods)
* [USE](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle#use)
* [TAG A VERSION](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle#tag-a-version)
* [UPDATE pkg.go.dev](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle#update-pkggodev)

Documentation and reference,

* Package at [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle)
* Refer to my example
  [module-with-remote-package](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/module-with-remote-package)

## TYPE & METHODS

```go
type Circle struct {
    R float64
}
```

```go
func (c Circle) **circleArea**() float64
func (c Circle) **circleCircumference**() float64
```

## USE

```bash
go get -u -v github.com/JeffDeCola/my-go-packages/geometry/circle
import github.com/JeffDeCola/my-go-packages/geometry/circle
```

## TAG A VERSION

To add a version to your package, you need to tag it before you commit
and push using the tag switch.

```bash
git add .
git tag v0.0.1
git commit -m "update circle"
git push --tags
```

## UPDATE pkg.go.dev

To publish a package on pkg.go.dev, head over to
[https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle)
and request it.
