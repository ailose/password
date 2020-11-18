# password
to generate password  automatically

### test
```
package main

import (
	"github.com/ailose/password"
	"log"
)


func main() {
	//get new password
	pwd := password.GenID()
	pwdInfo := password.GetNewPassword(pwd)
	// check password
	result, err := password.CheckPassword(pwdInfo.Pwd, pwdInfo.PasswordHash, pwdInfo.Salt)
	if err != nil {
		log.Fatal("check password err:", err)
	}
	if result == false {
		log.Fatal("the passowrd is wrong!")
	}
}

```
