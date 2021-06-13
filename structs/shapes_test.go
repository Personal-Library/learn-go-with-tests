package structs

import "testing"

// Rectangle has a method called Area that returns a float 64 so it passes
// Circle has a method called Area that also returns a float 64 so it passes
// In Go interface resolution is implicit. If the type you pass in matches
// what the interface is asking for, it will compile
type Shape interface {
	Area() float64
	Perimeter() float64
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 12, Width: 6}, want: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 62.83185307179586},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 28.970562748477136},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				// `%#v` format string will print out the struct with the values in its field
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
