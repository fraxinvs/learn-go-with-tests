package dependencyinjection

import (
	"bytes"
	"testing"
)

func Test_Greet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "John")

	got := buffer.String()
	want := "Hello, John"

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
