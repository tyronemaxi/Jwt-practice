package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	req "jwt-practice/api/request"
	"jwt-practice/api/response"
	"jwt-practice/service"
)

func UserRegister(c *gin.Context) {
	var userReq *req.UserInfoRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		response.Fail(c, err.Error())
		return
	}

	// check username is not exits
	userService := service.NewUserService()
	userCount, err := userService.GetUser(userReq.Username)
	if userCount > 0 {
		logrus.Errorf("用户名：[%s] 已存在， 请修改", userReq.Username)
		response.Fail(c, "用户名已存在，请修改")
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

	userInfo := service.CreateUserOption{
		UserInfo: userReq,
	}

	err = userService.CreateUser(userInfo)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(c, nil)

}

//func UserAuth(c *gin.Context) {
//	var token string
//	var err error
//
//	username, password := c.PostForm("username"), c.PostForm("password")
//
//	// step1: check the user exits
//
//	// step2: check the user password and username
//	if username == testUser.Username && password == testUser.Password {
//		token, err = util.GenerateJwtToken(username, password)
//		if err != nil {
//			c.JSON(http.StatusForbidden, gin.H{
//				"code": http.StatusForbidden,
//				"msg":  TokenGenerateFail,
//				"data": gin.H{
//					"token": nil,
//				},
//			})
//			return
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": http.StatusOK,
//		"msg":  "success",
//		"data": gin.H{
//			"token": token,
//		},
//	})
//}
