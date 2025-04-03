# MY PRIME PACKAGE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A package containing computations related to prime numbers._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-packages/tree/master/math/prime#overview)
  * [FUNCTIONS](https://github.com/JeffDeCola/my-go-packages/tree/master/math/prime#functions)
* [EXAMPLE](https://github.com/JeffDeCola/my-go-packages/tree/master/math/prime#example)  
* [ADD TO YOUR GO.MOD](https://github.com/JeffDeCola/my-go-packages/tree/master/math/prime#add-to-your-gomod)

Documentation and Reference

* Prime package at
  [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages/math/prime)

## OVERVIEW

This package contains a function to check if a number is prime.
It uses the Sieve of Eratosthenes algorithm to find all prime
numbers up to a given limit.
The Sieve of Eratosthenes is an ancient algorithm for finding all
prime numbers up to a specified integer.

### FUNCTIONS

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

## ADD TO YOUR GO.MOD

Since each package is tagged independently,

```text
git tag math/prime/vX.X.X
git push --tags
```

Add this to your go.mod file,

```text
require github.com/JeffDeCola/my-go-packages/math/prime vX.X.X
```
