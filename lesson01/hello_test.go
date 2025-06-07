package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHelloWithArgs(t *testing.T) {
	got := HelloWithArgs("Chris", "")
	want := "Hello, Chris!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHelloWithEmptyArgs(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloWithArgs("Chris", "")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when supplied with an empty string", func(t *testing.T) {
		got := HelloWithArgs("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := HelloWithArgs("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("without specifying a language", func(t * testing.T) {
		got := HelloWithArgs("Elodie", "")
		want := "Hello, Elodie!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := HelloWithArgs("Jules", "French")
		want := "Bonjour, Jules!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
