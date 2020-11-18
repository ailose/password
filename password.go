package password

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

const MinPasswordCost = 10

type PWD struct {
	Salt         string
	PasswordHash string
	Pwd          string
}

func GetNewPassword() *PWD {
	salt := genID2()
	pwd := genID()
	pwdMD5 := encodeMD5(pwd)
	passwordHash, _ := genPasswordHash(pwdMD5, salt, MinPasswordCost)
	return &PWD{
		Salt:         salt,
		Pwd:          pwd,
		PasswordHash: passwordHash,
	}
}

func CheckPassword(password, passwordHash, salt string) (bool, error) {
	pwd := encodeMD5(password)
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(salt+pwd)); err != nil {
		return false, err
	}
	return true, nil
}

func genPasswordHash(pwd string, salt string, cost int) (string, error) {
	bb, err := bcrypt.GenerateFromPassword([]byte(salt+pwd), cost)
	return string(bb), err
}
func genPassword() string {
	id := genID()
	return id[len(id)-8:]
}

func genID() string {
	return xid.New().String()
}

func genID2() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func encodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
