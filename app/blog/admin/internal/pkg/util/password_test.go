package util

import "testing"

func TestPasswordWithSalt(t *testing.T) {
	salt := PasswordWithSalt("1234");
	t.Log(" 加盐成功 ", salt)
}

func TestVerify(t *testing.T){
	verify := Verify("1234","$pbkdf2-sha512$5HsoqmjwO54KcPc4$c6ce24bf912ddae78ad4c8e5ed885e42f3595da868cd96777f764deebeb263e1")

	t.Log( verify )
}