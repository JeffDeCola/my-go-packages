// A package to calculate the area and circumference of a circle

package circle

import (
	"math"
)

//---------------------------------------------------------------
// Structs

// Circle struct defines a circle
type Circle struct {
	R float64
}

//---------------------------------------------------------------
// Methods

// CircleArea method to calculate the area of a circle
func (c Circle) CircleArea() float64 {
	return math.Pi * math.Pow(c.R, 2)
}

// CircleCircumference method to calculate the circumference of a circle
func (c Circle) CircleCircumference() float64 {
	return 2 * math.Pi * c.R
}
