package user

import (
	"gopkg.in/gomail.v2"
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
	secretPasswd, err := util.HashAndSalt(userOpt.UserInfo.Password)
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

func (u *UserService) GetUserCount(username string) (int, error) {
	userDao := user.NewUserDao()

	userCount, err := userDao.GetUserCount(user.UsernameFilter{
		UserName: []string{username},
	})

	return userCount, err
}

func (u *UserService) GetUserInfo(username string) (*user.User, error) {
	userDao := user.NewUserDao()

	userInfo, err := userDao.GetUser(username); if err != nil {
		return userInfo, err
	}

	return userInfo, nil
}

func userEmailVerify() {

}

func SendMail(mailTo []string, subject string, body string) error {
	// 设置邮箱主体
	mailConn := map[string]string{
		"user": "tyronemaxi@163.com",  //发送人邮箱（邮箱以自己的为准）
		"pass": "LULQUJFQFEIAAPPD",         //发送人邮箱的密码，现在可能会需要邮箱 开启授权密码后在pass填写授权码
		"host": "smtp.163.com", //邮箱服务器（此时用的是qq邮箱）
	}

	m := gomail.NewMessage(
		//发送文本时设置编码，防止乱码。 如果txt文本设置了之后还是乱码，那可以将原txt文本在保存时
		//就选择utf-8格式保存
		gomail.SetEncoding(gomail.Base64),
	)
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "邮箱验证")) // 添加别名
	m.SetHeader("To", mailTo...)                                  // 发送给用户(可以多个)
	m.SetHeader("Subject", subject)                               // 设置邮件主题
	m.SetBody("text/html", body)                                  // 设置邮件正文

	/*
	   创建SMTP客户端，连接到远程的邮件服务器，需要指定服务器地址、端口号、用户名、密码，如果端口号为465的话，
	   自动开启SSL，这个时候需要指定TLSConfig
	*/
	d := gomail.NewDialer(mailConn["host"], 465, mailConn["user"], mailConn["pass"]) // 设置邮件正文
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}

func NewUserService() *UserService {
	return &UserService{}
}
