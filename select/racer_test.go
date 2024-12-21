package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_Racer(t *testing.T) {
	slowServer := makeDeplayedServer(20 * time.Millisecond)
	fastServer := makeDeplayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func makeDeplayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}
