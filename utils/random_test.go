package utils

import (
	"strings"
	"testing"
)

func TestRandomString(t *testing.T) {
	s := RandomString()
	if len(strings.Split(s, "")) != 10 {
		t.Error("length is not 10")
		t.FailNow()
	}
	s = RandomString(14)
	if len(strings.Split(s, "")) != 14 {
		t.Error("length is not 14")
		t.FailNow()
	}
}
