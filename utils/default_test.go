package utils

import "testing"

func TestUnpackDefault(t *testing.T) {
	data := []string{"test", "test2"}
	defVal := "test3"
	if n := UnpackDefault(data, defVal); n != "test" {
		t.Error("unpack failed")
		t.FailNow()
	}
	if n := UnpackDefault([]string{}, defVal); n != "test3" {
		t.Error("unpack2 failed")
		t.FailNow()
	}
}
