package main

import (
	"fmt"
	"net/smtp"
	"strings"
)


type EmailClient struct {
	// ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerHost string
	// ServerPort 邮箱服务器端口，如腾讯企业邮箱为465，其它默认25
	ServerPort string
	// FromName　发件人邮箱名称，可以为空
	FromName string
	// FromEmail　发件人邮箱地址
	FromEmail string
	// FromPasswd 发件人邮箱客户端授权码（需要开通）注意：不是密码
	FromPasswd string
	// ToEmails 接收者邮件，如有多个，则以英文逗号(“,”)隔开，不能为空
	ToEmails []string
	// CcEmails 抄送者邮件，如有多个，则以英文逗号(“,”)隔开，可以为空
	CcEmails []string
}

func (ec *EmailClient) send(subject, body, mailType string) error {
	auth := smtp.PlainAuth("", ec.FromEmail, ec.FromPasswd, ec.ServerHost)
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	sendFrom := ec.FromEmail
	if ec.FromName != "" {
		sendFrom = ec.FromName + "<" + ec.FromEmail + ">"
	}

	var data []string
	data = append(data, "From: " + sendFrom)
	data = append(data, "To: " + strings.Join(ec.ToEmails, ","))
	//if ec.CcEmails != "" {
	//	data = append(data, "Cc: " + ec.CcEmails)
	//}
	data = append(data, "Subject: " + subject)
	data = append(data, contentType)
	data = append(data, "")
	data = append(data, body)
	msg := []byte(strings.Join(data, "\r\n"))
	serverAddress := strings.Join([]string{ec.ServerHost, ec.ServerPort}, ":")
	err := smtp.SendMail(serverAddress, auth, ec.FromEmail, ec.ToEmails, msg)
	return err
}

func main() {
	emailClient := EmailClient{
		ServerHost: "smtp.163.com",
		ServerPort: "25",
		FromName: "no-reply",
		FromEmail: "zhang_he06@163.com",
		FromPasswd: "xxxxxx",
		ToEmails: []string{"zhang_he06@163.com", "15555602203@163.com"},
	}
	subject := "使用Golang发送邮件"

	body := `
		<html>
		<body>
		<h3>
		Test send to email
		</h3>
		</body>
		</html>
		`
	err := emailClient.send(subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!" + err.Error())
	} else {
		fmt.Println("Send mail success!")
	}
}