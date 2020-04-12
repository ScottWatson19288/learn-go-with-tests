package shapes

import "math"

// Rectangle struct holds the width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct holds the radius
type Circle struct {
	Radius float64
}

// Shape interface that has the Area method
type Shape interface {
	Area() float64
}

// Perimeter takes in two float64 and returns the perimeter area
func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

// Area calculates the ares for a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area calculates the ares for a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
