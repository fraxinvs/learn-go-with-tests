package maps

import "testing"

func Test_Search(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("should return a known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("should return error for unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatalf("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error: %q, want: %q", got, want)
	}
}
