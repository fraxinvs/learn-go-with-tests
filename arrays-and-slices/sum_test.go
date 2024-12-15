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

		checkSums(t, got, want)
	})
}

func Test_SumAllTails(t *testing.T) {
	t.Run("should sum all items in collections except last item", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("should safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
