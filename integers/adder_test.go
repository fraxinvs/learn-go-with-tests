package integers

import (
	"fmt"
	"testing"
)

func Test_Adder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	assertCorrectAnswer(t, got, want)
}

func assertCorrectAnswer(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleAdd() {
	sum := Add(2, 4)
	fmt.Println(sum)
	// Output: 6
}
