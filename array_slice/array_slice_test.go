package array_slice

import (
	"fmt"
	"reflect"
	"testing"
)

// Main tests
func TestSum(t *testing.T) {
	assertArraySum := func(t testing.TB, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("\nGot: %d\nWant: %d\nNumbers: %v", got, want, numbers)
		}
	}

	t.Run("Array slice of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertArraySum(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	assertSumAll := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\nGot: %v\nWant: %v", got, want)
		}
	}

	t.Run("Adds multiple arrays in to an array of sums", func(t *testing.T){
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assertSumAll(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	assertSumAllTails := func (t testing.TB, got, want []int)  {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\nGot: %v\nWant: %v", got, want)
		}
	}

	t.Run("Adds all tails in arrays", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertSumAllTails(t, got, want)
	})

	t.Run("Safely sum empty arrays", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assertSumAllTails(t, got, want)
	})
}

// Examples for docs
func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func ExampleSumAll() {
	sumAll := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(sumAll)
	// Output: [3 9]
}

func ExampleSumAllTails() {
	summedTails := SumAllTails([]int{1, 2}, []int{0, 9})
	fmt.Println(summedTails)
	// Output: [2 9]
}

// Benchmarks
func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 2}, []int{0, 9})
	}
}