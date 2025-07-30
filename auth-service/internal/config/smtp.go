package config

import (
	"os"
	"strconv"
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
	portStr := os.Getenv("SMTP_PORT")
	port := 587 // default

	if portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			port = p
		}
	}

	SMTP = SMTPConfig{
		Host: os.Getenv("SMTP_HOST"),
		Port: port,
		User: os.Getenv("SMTP_USER"),
		Pass: os.Getenv("SMTP_PASS"),
		From: os.Getenv("SMTP_FROM"),
	}
}
