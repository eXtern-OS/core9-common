package utils

import "testing"

func TestValidateEmail(t *testing.T) {
	email1 := "test@test.com"
	email2 := "test test@.com"

	if !ValidateEmail(email1) {
		t.Error("error validating first email")
		t.FailNow()
	}

	if ValidateEmail(email2) {
		t.Error("error validating second email")
		t.FailNow()
	}
}
