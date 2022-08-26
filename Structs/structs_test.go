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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}

func assertMessageF(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}
