package token

import (
	"regexp"
	"testing"
)

var urlSafe = regexp.MustCompile(`^[A-Za-z0-9-_]+$`)

func TestEncode(t *testing.T) {
	for i := 0; i < 50; i++ {
		s, err := New(i)
		if err != nil {
			t.Fatal(err)
		}
		if !urlSafe.MatchString(s) {
			t.Errorf("%q contains characters that would need to be percent encoded in a url.", s)
		}
	}
}
