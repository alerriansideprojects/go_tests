package array_slice

import (
	"fmt"
	"testing"
)

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

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}