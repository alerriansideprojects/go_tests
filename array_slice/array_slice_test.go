package array_slice

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	assertArraySum := func(t testing.TB, got, want int, numbers [5]int) {
		t.Helper()
		if got != want {
			t.Errorf("\nGot: %d\nWant: %d\nNumbers: %v", got, want, numbers)
		}
	}

	t.Run("Add all numbers in array together", func(t *testing.T){
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assertArraySum(t, got, want, numbers)
	})
}

func ExampleSum() {
	numbers := [5]int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func BenchmarkSum(b *testing.B) {
	numbers := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}