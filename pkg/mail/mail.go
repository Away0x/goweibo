package mail

import (
	"crypto/tls"
	"errors"

	"github.com/lexkong/log"

	gomail "gopkg.in/gomail.v2"
)

const (
	// MailDriverSMTP -
	MailDriverSMTP = "smtp"
	// MailDriverLog -
	MailDriverLog = "log"
)

// Mail -
type Mail struct {
	Driver   string // smtp or log (log 时邮件是写在 log 中的，便于调试)
	Host     string // 邮箱的服务器地址
	Port     int    // 邮箱的服务器端口
	User     string // 发送者的 name
	Password string // 授权码或密码
	FromName string // 用来作为邮件的发送者名称

	MailTo []string // 发送目标

	Subject string // 邮件标题
	Body    string // 邮件内容
}

// Send 发送邮件
func (m *Mail) Send() error {
	if m.Driver == MailDriverLog {
		return m.sendByLog()
	} else if m.Driver == MailDriverSMTP {
		return m.sendBySMTP()
	}

	return errors.New("不支持该 Mail Driver: " + m.Driver)
}

func (m *Mail) sendBySMTP() error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.FromName+"<"+m.User+">")
	msg.SetHeader("To", m.MailTo...)
	msg.SetHeader("Subject", m.Subject)
	msg.SetBody("text/html", m.Body)

	d := gomail.NewDialer(m.Host, m.Port, m.User, m.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d.DialAndSend(msg)
}

func (m *Mail) sendByLog() error {
	log.Info(m.Body)
	return nil
}
