package sync

import "testing"

func Test_Counter(t *testing.T) {
	t.Run("should increment counter by 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
}

func assertCounter(t testing.TB, got Counter, want int) {
	if got.Value() != want {
		t.Errorf("got: %d, want: %d", got.Value(), 3)
	}
}
