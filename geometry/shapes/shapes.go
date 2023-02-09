// A package to calculate and manipulate simple 2d and 3d geometric shapes

package main

import (
	"math"
)

//---------------------------------------------------------------
// Interfaces

// TwoDCalculations interface of 2D shape calculations
type TwoDCalculations interface {
	Area() float64
	Perimeter() float64
}

// ThreeDCalculations interface of 3D shape calculations
type ThreeDCalculations interface {
	Volume() float64
	SurfaceArea() float64
}

// ShapeManipulations interface of 2D and 3D shape manipulations
type ShapeManipulations interface {
	Size(float64) // Pointer Receiver
}

//---------------------------------------------------------------
// Structs

// Rectangle struct defining it's geometry
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct defining it's geometry
type Circle struct {
	Radius float64
}

// Triangle struct defining it's geometry
type Triangle struct {
	A float64
	B float64
	C float64
}

// Cube struct defining it's geometry
type Cube struct {
	Edge float64
}

// Sphere struct defining it's geometry
type Sphere struct {
	Radius float64
}

//---------------------------------------------------------------
// Methods

// Area method to calculate the area of a rectangle
func (r Rectangle) Area() float64 {
	a := r.Width * r.Height
	return a
}

// Area method to calculate the area of a circle
func (c Circle) Area() float64 {
	a := math.Pi * math.Pow(c.Radius, 2)
	return a
}

// Area method to calculate the area of a triangle
func (t Triangle) Area() float64 {
	// Heron's Formula to get area from 3 sides
	s := ((t.A + t.B + t.C) / 2)
	a := math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
	return a
}

// Perimeter method to calculate the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	p := 2 * (r.Width + r.Height)
	return p
}

// Perimeter method to calculate the perimeter of a circle
func (c Circle) Perimeter() float64 {
	p := 2 * math.Pi * c.Radius
	return p
}

// Perimeter method to calculate the perimeter of a triangle
func (t Triangle) Perimeter() float64 {
	p := t.A + t.B + t.C
	return p
}

// Volume method to calculate the volume of a cube
func (c Cube) Volume() float64 {
	v := c.Edge * c.Edge * c.Edge
	return v
}

// Volume method to calculate the volume of a sphere
func (s Sphere) Volume() float64 {
	v := (4.0 / 3.0) * math.Pi * math.Pow(s.Radius, 3)
	return v
}

// SurfaceArea method to calculate the surface area of a cube
func (c Cube) SurfaceArea() float64 {
	sa := 6 * c.Edge * c.Edge
	return sa
}

// SurfaceArea method to calculate the surface area of a sphere
func (s Sphere) SurfaceArea() float64 {
	sa := 4 * math.Pi * math.Pow(s.Radius, 2)
	return sa
}

// Size method to scale a rectangle
func (r *Rectangle) Size(f float64) {
	r.Width = f * r.Width
	r.Height = f * r.Height
}

// Size method to scale a circle
func (c *Circle) Size(f float64) {
	c.Radius = f * c.Radius
}

// Size method to scale a triangle
func (t *Triangle) Size(f float64) {
	t.A = f * t.A
	t.B = f * t.B
	t.C = f * t.C
}

// Size method to scale a cube
func (c *Cube) Size(f float64) {
	c.Edge = f * c.Edge
}

// Size method to scale a sphere
func (s *Sphere) Size(f float64) {
	s.Radius = f * s.Radius
}

//---------------------------------------------------------------
// Functions

// GetArea function returns the area of a 2D shape
func GetArea(t TwoDCalculations) float64 {
	return t.Area()
}

// GetPerimeter function returns the perimeter of a shape
func GetPerimeter(t TwoDCalculations) float64 {
	return t.Perimeter()
}

// GetVolume function returns the volume of a 3D shape
func GetVolume(t ThreeDCalculations) float64 {
	return t.Volume()
}

// GetSurfaceArea function the perimeter of a shape
func GetSurfaceArea(t ThreeDCalculations) float64 {
	return t.SurfaceArea()
}

// ChangeSize function changes the size of a shape
func ChangeSize(t ShapeManipulations, f float64) {
	t.Size(f)
}
