package utils

import (
	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(mailTo string, subject string, body string) error {
	mailConn := map[string]string{
		"user": beego.AppConfig.String("mail_user"),
		"pass": beego.AppConfig.String("mail_pass"),
		"host": "smtp.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"])

	m := gomail.NewMessage()

	m.SetHeader("From",  m.FormatAddress(mailConn["user"], "Beego-user"))
	m.SetHeader("To", mailTo)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err
}