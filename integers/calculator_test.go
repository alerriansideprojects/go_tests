package integers

import (
	"fmt"
	"testing"
)

func TestCalculator(t *testing.T) {
	assertMathFunction := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("\nGot: '%d'\nWant: '%d'", got, want)
		}
	}

	t.Run("Add 2 integers", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assertMathFunction(t, got, want)
	})

	t.Run("Subtract 2 integers", func(t *testing.T) {
		got := Subtract(2, 1)
		want := 1
		assertMathFunction(t, got, want)
	})

	t.Run("Multiply 2 integers", func(t *testing.T) {
		got := Multiply(3, 5)
		want := 15
		assertMathFunction(t, got, want)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func ExampleSubtract() {
	difference := Subtract(2, 1)
	fmt.Println(difference)
	// Output: 1
}

func ExampleMultiply() {
	product := Multiply(3, 5)
	fmt.Println(product)
	// Output: 15
}