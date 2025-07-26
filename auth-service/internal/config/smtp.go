// config/smtp.go
package config

import (
	"os"
)

type SMTPConfig struct {
	Host string
	Port int
	User string
	Pass string
	From string
}

var SMTP SMTPConfig

func InitSMTP() {
	// Convert port ke int
	port := 587 // default
	SMTP = SMTPConfig{
		Host: os.Getenv("SMTP_HOST"),
		Port: port,
		User: os.Getenv("SMTP_USER"),
		Pass: os.Getenv("SMTP_PASS"),
		From: os.Getenv("SMTP_FROM"),
	}
}
