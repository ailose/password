package password

import (
	"testing"
)

func TestGenID(t *testing.T) {
	//get new password
	pwd := GenID()
	_, err := GetNewPassword(pwd)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetNewPassword(t *testing.T) {
	//get new password
	pwd := GenID()
	_, err := GetNewPassword(pwd)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckPassword(t *testing.T) {
	//get new password
	pwd := GenID()
	pwdInfo, err := GetNewPassword(pwd)
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
