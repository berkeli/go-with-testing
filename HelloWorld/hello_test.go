package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Berkeli", "English")
		want := "Hello, Berkeli!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello world on empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in Spanish", func(t *testing.T) {
		got := Hello("Berkeli", "Spanish")
		want := "Hola, Berkeli!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in French", func(t *testing.T) {
		got := Hello("Berkeli", "French")
		want := "Hola, Berkeli!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
