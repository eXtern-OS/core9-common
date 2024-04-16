package utils

import (
	"fmt"
	"testing"
)

func TestSHA256(t *testing.T) {
	data := "hello"
	expected := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	res := SHA256(data)
	if res != expected {
		t.Error(fmt.Sprintf("hashes don't match:%s\n%s\n", data, expected))
	}
}
