package email

import (
	"gopkg.in/gomail.v2"
)

const (
	// 邮件服务器地址
	MAIL_HOST = "smtp.qq.com"
	// 端口
	MAIL_PORT = 465
	// 发送邮件用户账号
	MAIL_USER = "1476580586@qq.com"
	// 授权密码
	MAIL_PWD = "keprlhmbnbmlgbbe"
)

/*
title 使用gomail发送邮件
@param []string mailAddress 收件人邮箱
@param string subject 邮件主题
@param string body 邮件内容
@return error
*/
func SendGoMail(mailAddress []string, subject string, body string) error {
	m := gomail.NewMessage()
	// 这种方式可以添加别名，即 nickname， 也可以直接用<code>m.SetHeader("From", MAIL_USER)</code>
	nickname := "go-blog"
	m.SetHeader("From", nickname+"<"+MAIL_USER+">")
	// 发送给多个用户
	m.SetHeader("To", mailAddress...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(MAIL_HOST, MAIL_PORT, MAIL_USER, MAIL_PWD)
	// 发送邮件
	err := d.DialAndSend(m)

	return err
}
