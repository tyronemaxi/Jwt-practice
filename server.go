package main
//
//import (
//	"fmt"
//	"gopkg.in/gomail.v2"
//)
//
///*
//go邮件发送
//*/
//
//
//
//func main() {
//	// 邮件接收方
//	mailTo := []string{
//		//可以是多个接收人
//		//"493537558@qq.com",
//		"tyronextian@gmail.com",
//	}
//
//	subject := "验证你的电子邮件!" // 邮件主题
//	verify_code := 123456
//	body := fmt.Sprintf("验证你是该电子邮箱: [xxxx@qq.com] 的拥有者，邮箱验证码为： [%d]", verify_code)        // 邮件正文
//
//	err := SendMail(mailTo, subject, body)
//	if err != nil {
//		fmt.Println("Send fail! - ", err)
//		return
//	}
//	fmt.Println("Send successfully!")
//}*/