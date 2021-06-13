package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// It is convention in Go to have the receiver be the first letter of the type
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
