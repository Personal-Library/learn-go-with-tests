package main

import "testing"

func TestHello(t *testing.T) {
	// testing.TB is an interface which *testing.T and *testing.B both satisfy,
	// so you can call helper functions from a test or a benchmark
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Anthony", "")
		want := "Hello, Anthony"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(*testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola, Juan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Amelie", "French")
		want := "Bonjour, Amelie"
		assertCorrectMessage(t, got, want)
	})
}
