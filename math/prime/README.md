# MY PRIME PACKAGE

_A package containing computations related to prime numbers._

Table of Contents

* [FUNCTIONS](https://github.com/JeffDeCola/my-go-packages/tree/master/math/prime#functions)
* [EXAMPLE](https://github.com/JeffDeCola/my-go-packages/tree/master/math/prime#example)

Documentation and Reference

* Prime package at
  [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/math/prime)

## FUNCTIONS

```go
func IsPrime(n int) bool
```

## EXAMPLE

```go
package main

import (
    "fmt"

    "github.com/JeffDeCola/my-go-packages/math/prime"
)

func main() {

    // Check if a number is prime
    fmt.Println("Is 7 prime?", prime.IsPrime(7))
    fmt.Println("Is 8 prime?", prime.IsPrime(8))

}
```

## ADDED TO YOUR GO.MOD

Since I am tagging each package independently,

```text
git tag math/prime/vX.X.X
git push --tags
```

This will be added to your go.mod file,

```text
require github.com/JeffDeCola/my-go-packages/math/prime vX.X.X
```
