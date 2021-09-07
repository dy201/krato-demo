package util

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"strings"
)

// 密码加盐值
func PasswordWithSalt(pwd string) string {
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(pwd, options)
	return fmt.Sprintf("$pbkdf2-sha512$%s$%s",salt,encodedPwd)
}

func Verify(pwd string, encodedPwdInfo string) bool{
	//encodedPwd := "$pbkdf2-sha512$5HsoqmjwO54KcPc4$c6ce24bf912ddae78ad4c8e5ed885e42f3595da868cd96777f764deebeb263e1"
	params := strings.Split(encodedPwdInfo, "$")
	salt := params[2]
	encodedPwd := params[3]
	options := &password.Options{16, 100, 32, sha512.New}
	return password.Verify(pwd,salt,encodedPwd,options)
}
