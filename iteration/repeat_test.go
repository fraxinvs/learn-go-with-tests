package iteration

import (
	"fmt"
	"testing"
)

func Test_Repeat(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func Benchmark_Repeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}
