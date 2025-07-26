package utils

import (
	"auth-service/internal/config"

	"gopkg.in/gomail.v2"
)

func SendMail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.SMTP.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(
		config.SMTP.Host,
		config.SMTP.Port,
		config.SMTP.User,
		config.SMTP.Pass,
	)

	return d.DialAndSend(m)
}
