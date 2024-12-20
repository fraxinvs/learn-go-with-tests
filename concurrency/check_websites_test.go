package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "invalid://test-invalid-website.com"
}

func Test_CheckWebsites(t *testing.T) {
	websites := []string{
		"http://test-valid-website.com",
		"http://test-valid-website-2.com",
		"invalid://test-invalid-website.com",
	}

	want := map[string]bool{
		"http://test-valid-website.com":      true,
		"http://test-valid-website-2.com":    true,
		"invalid://test-invalid-website.com": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted: %v, got: %v", want, got)
	}
}
