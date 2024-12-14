package main

import "testing"

func Test_Hello(t *testing.T) {
	t.Run("should reply hello to given person", func(t *testing.T) {
		got := Hello("John", "English")
		want := "Hello, John"

		assertCorrectMessage(t, got, want)
	})

	t.Run("should reply 'Hello, World' when given an empty string", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("should reply with Spanish greeting", func(t *testing.T) {
		got := Hello("Marc", "Spanish")
		want := "Hola, Marc"

		assertCorrectMessage(t, got, want)
	})

	t.Run("should reply with French greeting", func(t *testing.T) {
		got := Hello("Jules", "French")
		want := "Bonjour, Jules"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
