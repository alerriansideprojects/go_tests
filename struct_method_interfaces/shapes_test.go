package struct_method_interfaces

import (
	"fmt"
	"testing"
)

// Main tests
func TestPerimeter(t *testing.T) {
	assertPerimeter := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("\nGot: %.2f\nWant: %.2f", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		assertPerimeter(t, got, want)
	})
}

func TestArea(t *testing.T) {
	assertArea := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("\nGot: %g\nWant: %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		assertArea(t, got, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		assertArea(t, got, want)
	})
}

// Examples
func ExamplePerimeter() {
	rectangle := Rectangle{10.0, 10.0}
	perimeter := rectangle.Perimeter()
	fmt.Println(perimeter)
	// Output: 40
}

func ExampleArea() {
	rectangle := Rectangle{10.0, 10.0}
	rec_area := rectangle.Area()
	fmt.Println(rec_area)
	// Output: 100
}

// Benchmarks
func BenchmarkRecPerimeter(b *testing.B) {
	rectangle := Rectangle{10.0, 10.0}
	for i := 0; i < b.N; i++ {
		rectangle.Perimeter()
	}
}

func BenchmarkRecArea(b *testing.B) {
	rectangle := Rectangle{10.0, 10.0}
	for i := 0; i < b.N; i++ {
		rectangle.Area()
	}
}

func BenchmarkCirArea(b *testing.B) {
	circle := Circle{10}
	for i := 0; i < b.N; i++ {
		circle.Area()
	}
}