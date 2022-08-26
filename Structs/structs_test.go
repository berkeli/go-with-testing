package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Perimeter of 10x10", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := Perimeter(rect)
		want := 40.0
		assertMessageF(t, got, want)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}

	t.Run("Area of 12x6 Rectangle", func(t *testing.T) {
		rect := Rectangle{12.0, 6.0}
		checkArea(t, rect, 72.0)
	})

	t.Run("area of a circle", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}

func assertMessageF(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}
