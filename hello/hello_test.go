package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("\nGot: %q\nWant: %q", got, want)
		}
	}

	t.Run("Saying hello to specific people", func(t *testing.T){
		got := Hello("Steve!", "")
		want := "Hello, Steve!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, Go!' when an empty string is supplied", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, Go!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Message in Spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Message in French", func(t *testing.T){
		got := Hello("Anya", "French")
		want := "Bonjour, Anya"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Message in German", func(t *testing.T){
		got := Hello("Karl", "German")
		want := "Hallo, Karl"

		assertCorrectMessage(t, got, want)
	})
}