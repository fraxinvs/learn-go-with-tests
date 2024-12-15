package arraysandslices

import (
	"slices"
	"testing"
)

func Test_Sum(t *testing.T) {
	t.Run("should sum a collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got: %q, want: %q, given: %v", got, want, numbers)
		}
	})
}

func Test_SumAll(t *testing.T) {
	t.Run("should sum two slices and return sum of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
