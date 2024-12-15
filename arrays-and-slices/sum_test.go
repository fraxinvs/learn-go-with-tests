package arraysandslices

import "testing"

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
