package main

import "testing"

func Test_Hello(t *testing.T) {
	t.Run("should reply hello to given person", func(t *testing.T) {
		got := Hello("John")
		want := "Hello, John"

		if got != want {
			t.Errorf("got %q want %q ", got, want)
		}
	})

	t.Run("should reply 'Hello, World' when given an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q ", got, want)
		}
	})
}
