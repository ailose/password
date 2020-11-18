package password

import (
	"testing"
)

func TestGetNewPassword(t *testing.T) {
	//get new password
	_, err := GetNewPassword()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckPassword(t *testing.T) {
	//get new password
	pwdInfo, err := GetNewPassword()
	if err != nil {
		t.Fatal(err)
	}
	// check password
	result, err := CheckPassword(pwdInfo.Pwd, pwdInfo.PasswordHash, pwdInfo.Salt)
	if err != nil {
		t.Error("check password err:", err)
	}
	if result == false {
		t.Error("the passowrd is wrong!")
	}
}
