package request

import (
	"testing"
)

func TestUserRegistered_UserRegisterValidateMatch(t *testing.T) {
	var err error
	a := map[string]string{
		"username":     "Allen",      // err name
		"password":     "Allen",      // err password
		"mobile_phone": "173626",     // err mobile
		"email":        "123411.com", // err email
	}

	b := map[string]string{
		"username":     "Allen",
		"password":     "tyrone@1234!",
		"mobile_phone": "13041479216",
		"email":        "tyronemaxi@163.com",
	}

	userValidate := NewUserAuthValidate()
	err = userValidate.UserRegisterValidate(a)
	if err != nil {
		t.Errorf("validate err: [%s]", err.Error())
	}

	err = userValidate.UserRegisterValidate(b)
	if err != nil {
		t.Errorf("validate err: [%s]", err.Error())
	}

}
