package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"website/api"
	req "website/api/request"
	"website/api/response"
	"website/service/user"
	"website/util"
)

func Register(c *gin.Context) {
	var userReq *req.UserInfoRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		response.Fail(c, err.Error())
		return
	}

	// check username is not exits
	userService := user.NewUserService()
	userCount, err := userService.GetUserCount(userReq.Username)
	if userCount > 0 {
		logrus.Errorf("用户名：[%s] 已存在， 请修改", userReq.Username)
		response.Fail(c, "用户名已存在，请重新输入")
		return
	}

	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	validateUserOption := map[string]string{
		"username":     userReq.Username,
		"password":     userReq.Password,
		"mobile_phone": userReq.MobilePhone,
		"email":        userReq.Email,
	}

	// validate
	logrus.Infof("validate the UserRegister option for user %s", userReq.Username)
	userValidate := req.NewUserAuthValidate()
	err = userValidate.UserRegisterValidate(validateUserOption)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	logrus.Infof("Create the new user: %s", userReq.Username)


	userInfo := user.CreateUserOption{
		UserInfo: userReq,
	}

	err = userService.CreateUser(userInfo)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	// 验证邮箱或者手机号

	response.Success(c, nil)

}

func Auth(c *gin.Context) {
	var err error
	var token string
	username, password := c.PostForm("username"), c.PostForm("password")

	userService := user.NewUserService()
	user, err := userService.GetUserInfo(username); if err != nil {
		logrus.Errorf("用户名：[%s] 不存在， 请重新输入",username)
		response.Fail(c, "用户名不存在， 请重新输入")
		return
	}

	if util.ComparePasswords(user.Password, password) {
		token, err = util.GenerateJwtToken(username, password)
		if err != nil {
			response.Fail(c, api.TokenGenerateFail)
			return
		}
	} else {
		logrus.Errorf("用户密码：%s 错误", password)
		response.Fail(c, "用户密码错误，请重新输入")
		return
	}

	response.Success(c, token)
}

//
//func VerifyMobile(c *gin.Context) {
//
//}
//
//func VerifyEmail(c *gin.Context) {
//	var emailTocken string
//	var err error
//	var emailMsg *req.UserEmailVerify
//
//	if err := c.ShouldBindJSON(&emailMsg); err != nil {
//		response.Fail(c, err.Error())
//		return
//	}
//
//
//	// 解析 user 信息
//
//	// 查询用户信息 uid
//
//	// 根据 uid, email, 生成 uid
//
//	if util.ComparePasswords(emailMsg.Email,  ) {
//		emailTocken, err = util.GenerateJwtToken(username, password)
//		if err != nil {
//			response.Fail(c, api.TokenGenerateFail)
//			return
//		}
//	} else {
//		logrus.Errorf("用户密码：%s 错误", password)
//		response.Fail(c, "用户密码错误，请重新输入")
//		return
//	}





//}
