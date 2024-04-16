package user

import "testing"

func TestUser_Verify(t *testing.T) {
	u1 := User{Login: "test_login"}
	u2 := User{Login: "test login $"}

	if err := u1.Verify(); err != nil {
		t.Error(err)
		t.FailNow()
	}

	if err := u2.Verify(); err == nil {
		t.Error("matched when should have failed")
		t.FailNow()
	}
}
