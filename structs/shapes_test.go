package structs

import "testing"

// Rectangle has a method called Area that returns a float 64 so it passes
// Circle has a method called Area that also returns a float 64 so it passes
// In Go interface resolution is implicit. If the type you pass in matches
// what the interface is asking for, it will compile
type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	// .2 means print 2 decimal places
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		// Using g will print a more precise deicmal number in the error message
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{2.0, 4.0}
		checkArea(t, rectangle, 8.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
