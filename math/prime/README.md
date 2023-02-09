# My prime package

_A package containing computations related to prime numbers._

Documentation and reference,

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

Where go.mod is,

```text
module your-module-name

go 1.19

require github.com/JeffDeCola/my-go-packages v0.2.0
```
