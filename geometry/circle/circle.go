package circle

// A package to calculate the area and circumference of a circle

import (
	"math"
)

//---------------------------------------------------------------
// Structs

// Circle struct defines a circle
type Circle struct {
	Radius float64
}

//---------------------------------------------------------------
// Methods

// Area method to calculate the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Circumference method to calculate the circumference of a circle
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}
