package config

import (
	"net/smtp"
	"os"
)

type SMTPMailConfig struct {
	Host          string
	Port          string
	Username      string
	Password      string
	SenderName    string
	SenderAddress string
}

func NewSMTPConfiguration() SMTPMailConfig {
	return SMTPMailConfig{
		Host:          os.Getenv("MAIL_HOST"),
		Port:          os.Getenv("MAIL_PORT"),
		Username:      os.Getenv("MAIL_USERNAME"),
		Password:      os.Getenv("MAIL_PASSWORD"),
		SenderName:    os.Getenv("MAIL_FROM_NAME"),
		SenderAddress: os.Getenv("MAIL_FROM_ADDRESS"),
	}
}

func (c SMTPMailConfig) SMTPAuth() smtp.Auth {
	return smtp.PlainAuth("", c.Username, c.Password, c.Host)
}
