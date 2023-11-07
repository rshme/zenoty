package usecase

import (
	"bytes"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/rafiseptian90/zenoty/internal/example_app/config"
	"github.com/rafiseptian90/zenoty/internal/example_app/domain"
)

type ReminderEmailNotification struct {
	config config.SMTPMailConfig
}

func NewReminderEmailNotification(mailConfig config.SMTPMailConfig) *ReminderEmailNotification {
	return &ReminderEmailNotification{config: mailConfig}
}

func (r *ReminderEmailNotification) SendEmail(receivers []string, subject string, body string) error {
	smtpAuth := r.config.SMTPAuth()
	smtpAddr := fmt.Sprintf("%s:%s", r.config.Host, r.config.Port)

	message := r.GenerateMessageBuffer(
		mail.Address{
			Name:    r.config.SenderName,
			Address: r.config.SenderAddress,
		},
		receivers[0],
		receivers[1:],
		subject,
		body,
	)

	if err := smtp.SendMail(smtpAddr, smtpAuth, r.config.SenderAddress, receivers, message.Bytes()); err != nil {
		return err
	}

	return nil
}

func (r *ReminderEmailNotification) PrepareHTMLTemplate(reminder domain.Reminder) (string, error) {
	htmlPath := fmt.Sprint("templates/reminder.html")

	tmpl, err := template.ParseFiles(htmlPath)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	if err := tmpl.Execute(&result, reminder); err != nil {
		return "", err
	}

	return result.String(), nil
}

func (r *ReminderEmailNotification) GenerateMessageBuffer(fromAddress mail.Address, to string, cc []string, subject string, body string) bytes.Buffer {
	headers := make(map[string]string)
	headers["From"] = fromAddress.String()
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""

	if len(cc) > 0 {
		headers["Cc"] = strings.Join(cc, ",")
	}

	messageBytes := []byte(body)

	// Combine the message headers and body into a single message.
	var buffer bytes.Buffer
	for k, v := range headers {
		buffer.WriteString(k)
		buffer.WriteString(": ")
		buffer.WriteString(v)
		buffer.WriteString("\r\n")
	}
	buffer.WriteString("\r\n")
	buffer.Write(messageBytes)

	return buffer
}
