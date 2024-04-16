package utils

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	data := "hello"
	encoded := EncodeBase64(data)
	decoded, err := DecodeBase64(encoded)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if decoded != data {
		t.Error(fmt.Sprintf("failed, strings do not match: %s != %s", data, decoded))
	}
}
