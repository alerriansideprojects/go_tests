package maps

import "testing"

func TestSearch(t *testing.T) {
	assertStrings := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\nGot: %q\nWant: %q", got, want)
		}
	}

	t.Run("Known word", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is just a test"}

		got := Search(dictionary, "test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
}