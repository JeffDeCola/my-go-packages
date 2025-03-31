# MY GO PACKAGES

[![Go Reference](https://pkg.go.dev/badge/github.com/JeffDeCola/my-go-packages.svg)](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages)
[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/my-go-packages)](https://goreportcard.com/report/github.com/JeffDeCola/my-go-packages)
[![codeclimate Maintainability](https://api.codeclimate.com/v1/badges/429352c4ab8e00602452/maintainability)](https://codeclimate.com/github/JeffDeCola/my-go-packages/maintainability)
[![codeclimate Issue Count](https://codeclimate.com/github/JeffDeCola/my-go-packages/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/my-go-packages/issues)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)
[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)

_A place to keep my go packages._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-packages#overview)
* [GEOMETRY](https://github.com/JeffDeCola/my-go-packages#geometry)
* [GOLANG](https://github.com/JeffDeCola/my-go-packages#golang)
* [MATH](https://github.com/JeffDeCola/my-go-packages#math)
* [NEURAL NETWORKS](https://github.com/JeffDeCola/my-go-packages#neural-networks)

Documentation and Reference

* [pkg.go.dev](https://pkg.go.dev/github.com/JeffDeCola/my-go-packages)
  shows these packages
* [go-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet)
* [my-go-examples](https://github.com/JeffDeCola/my-go-examples)
* [my-go-tools](https://github.com/JeffDeCola/my-go-tools)
* This repos
  [github webpage](https://jeffdecola.github.io/my-go-packages/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-packages/blob/master/ci-README.md)_

## OVERVIEW

Every package is tagged with it's own version and has it's
own go.mod file. This is done to prevent downloading the entire repo for
your dependencies. You only get what you want.
For example, if you want the circle package, your go.mod would look like,

```text
require github.com/JeffDeCola/my-go-packages/geometry/circle v1.0.0
```

## GEOMETRY

* [![Tag Latest](https://img.shields.io/badge/v0.0.1-blue)](https://github.com/JeffDeCola/my-go-packages/releases/tag/geometry/circle/v0.0.1)
  [circle](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/circle)

  _A package to calculate the area and circumference of a circle._

* [![Tag Latest](https://img.shields.io/badge/v0.0.1-blue)](https://github.com/JeffDeCola/my-go-packages/releases/tag/geometry/shapes/v0.0.1)
  [shapes](https://github.com/JeffDeCola/my-go-packages/tree/master/geometry/shapes)

  _A package to calculate and manipulate simple 2D and 3D geometric shapes._

## GOLANG

* [![Tag Latest](https://img.shields.io/badge/v0.1.0-blue)](https://github.com/JeffDeCola/my-go-packages/releases/tag/golang/logger/v0.1.0)
  [logger](https://github.com/JeffDeCola/my-go-packages/tree/master/golang/logger)

  _Just a logger wrapper formatting a message followed
  by key-value pairs.
  Currently using the standard library
  [slog](https://pkg.go.dev/log/slog)
  supporting both text and json._

## MATH

* [![Tag Latest](https://img.shields.io/badge/v0.0.1-blue)](https://github.com/JeffDeCola/my-go-packages/releases/tag/math/prime/v0.0.1)
  [prime](https://github.com/JeffDeCola/my-go-packages/tree/master/math/prime)

  _A package containing computations related to prime numbers._

## NEURAL NETWORKS

* [![Tag Latest](https://img.shields.io/badge/v0.0.1-blue)](https://github.com/JeffDeCola/my-go-packages/releases/tag/neural-networks/mlp/v0.0.1)
  [mlp](https://github.com/JeffDeCola/my-go-packages/tree/master/neural-networks/mlp)

  _A package to implement a scalable multi-layer
  perceptron (MLP) neural network._
