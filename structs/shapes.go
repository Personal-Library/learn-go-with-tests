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

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return (2 * c.Radius) * math.Pi
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) Perimeter() float64 {
	// ASSUMING AN EQUILATERAL TRIANGLE
	// Pythagorean Theorem a^2 + b^2 = c^2
	a := t.Base / 2
	b := t.Height
	cSquared := (a * a) + (b * b)
	c := math.Sqrt(cSquared)
	return t.Base + c + c
}
