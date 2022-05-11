package service

import (
	"golang.org/x/crypto/bcrypt"
	"jwt-practice/api/request"
	"jwt-practice/database/user"
	"jwt-practice/util"
)

const (
	UserCreateType = "Create"
	UserImportType = "Import"
)

type CreateUserOption struct {
	UserInfo       *request.UserInfoRequest
	UserID         string `json:"user_id"`
	DomainID       string `json:"domain_id"`
	ProjectID      string `json:"project_id"`
	DisplayName    string `json:"display_name"`
	SourceType     string `json:"source_type"`
	MobileVerified int    `json:"mobile_verified"`
	EmailVerified  string `json:"email_verified"`
	Extra          string `json:"extra"`
}

type UserService struct {
}

func (u *UserService) CreateUser(userOpt CreateUserOption) error {
	// 用户名不唯一
	userID, err := util.GenerateUUid("user")
	if err != nil {
		return err
	}
	// 别名不唯一

	// 用户名密码 brcypt 加密
	pwd := []byte(userOpt.UserInfo.Password)
	secretPasswd, err := bcrypt.GenerateFromPassword(pwd, 10)
	if err != nil {
		return err
	}

	// displayN
	if userOpt.DisplayName == "" {
		// auto generate
		userOpt.DisplayName = userOpt.UserInfo.Username
	}

	if userOpt.SourceType == "" {
		userOpt.SourceType = UserCreateType
	}

	CurrentUser := user.User{
		UserName:       userOpt.UserInfo.Username,
		DisplayName:    userOpt.DisplayName,
		UserID:         userID,
		Password:       string(secretPasswd),
		MobilePhone:    userOpt.UserInfo.MobilePhone,
		MobileVerified: 0,
		Email:          userOpt.UserInfo.Email,
		EmailVerified:  0,
	}

	userDao := user.NewUserDao()
	err = userDao.CreateUser(CurrentUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetUser(username string) (int, error) {
	userDao := user.NewUserDao()

	userCount, err := userDao.GetUserCount(user.UsernameFilter{
		UserName: []string{username},
	})

	return userCount, err
}

func NewUserService() *UserService {
	return &UserService{}
}
