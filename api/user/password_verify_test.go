package user

import (
	"fmt"
	"jwt-practice/util"
	"testing"
)

func TestPasswordHashAndCompare(t *testing.T) {
	var err error
	passwordOK := "admin"
	//passwordERR := "adminxx"

	hashStr, err := util.HashAndSalt(passwordOK)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hashStr)

	// 正确密码验证
	check := util.ComparePasswords(hashStr, passwordOK)
	if !check {
		t.Errorf("pw wrong")

	}
}
