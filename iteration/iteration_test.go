package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	assertRepeatFunction := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\nGot: %q\nWant: %q", got, want)
		}
	}

	t.Run("Repeats a letter", func(t *testing.T) {
		got := Repeat("a", 1)
		want := "a"
		assertRepeatFunction(t, got, want)
	})

	t.Run("Repeats as much as the user wants", func(t *testing.T) {
		got := Repeat("s", 10)
		want := "ssssssssss"
		assertRepeatFunction(t, got, want)
	})
}

func ExampleRepeat() {
	repeated := Repeat("s", 10)
	fmt.Println(repeated)
	// Output: ssssssssss
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}