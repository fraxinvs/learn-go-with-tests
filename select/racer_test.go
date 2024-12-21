package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_Racer(t *testing.T) {
	slowServer := makeDeplayedServer(20 * time.Nanosecond)
	fastServer := makeDeplayedServer(0 * time.Nanosecond)

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}

func makeDeplayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}
