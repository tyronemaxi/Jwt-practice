package request

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"regexp"
)

type UserInfoRequest struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	MobilePhone    string `json:"mobile_phone" binding:"required"`
	MobileVerified int    `json:"mobile_verified"`
	Email          string `json:"email" binding:"required"`
	EmailVerified  int    `json:"email_verified"`
}

type UserEmailVerify struct {
	Email        string `json:"email" binding:"required"`
	SendIfExits  bool   `json:"send_if_exits" binding:"required"`
	UserMsgToken string `json:"user_token" binding:"requeired"`
}

const (
	UserNameRule    = "^[\u4e00-\u9fa5a-zA-Z0-9]{5,20}$"
	PasswordRule    = "^([A-Z]|[a-z]|[0-9]|[-=[;,./~!@#$%^*()_+}{:?]){6,20}$"
	MobilePhoneRule = "^1(3\\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\\d|9[0-35-9])\\d{8}$"
	EmailRule       = "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
)

func userRegisterValidateInit(userField map[string]string) (map[string]*regexp.Regexp, error) {
	userRegexpRule := map[string]string{
		"username":     UserNameRule,
		"password":     PasswordRule,
		"mobile_phone": MobilePhoneRule,
		"email":        EmailRule,
	}
	ruleMap := make(map[string]*regexp.Regexp, len(userField))

	var err error
	for k, v := range userRegexpRule {
		ruleMap[k], err = regexp.Compile(v)
		if err != nil {
			return nil, err
		}
	}
	return ruleMap, nil
}

func (u *UserInfoRequest) UserRegisterValidate(userInfo map[string]string) error {
	if len(userInfo) == 0 {
		return fmt.Errorf("the validateFiled is empty")
	}

	reRule, err := userRegisterValidateInit(userInfo)
	if err != nil {
		return err
	}

	for userFiled, userValue := range userInfo {
		logrus.Infof("validate the UserRegister option key:[%s], value:[%s]", userFiled, userValue)
		for k, v := range reRule {
			if k == userFiled {
				if !v.MatchString(userValue) {
					logrus.Errorf("user auth failed, filed: [%s], value: [%s]\n", userFiled, userValue)
					return fmt.Errorf("user auth failed, filed: [%s], value: [%s]", userFiled, userValue)
				}
			}
		}
	}

	return nil
}

func NewUserAuthValidate() *UserInfoRequest {
	return &UserInfoRequest{}
}
