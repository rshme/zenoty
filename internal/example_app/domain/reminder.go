package domain

import (
	"bytes"
	"net/mail"
)

type Task struct {
	TaskTitle       string `json:"task_title"`
	TaskDescription string `json:"task_description"`
	TaskDate        string `json:"task_date"`
}

type Reminder struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	Token string `json:"token,omitempty"`
	Tasks []Task `json:"tasks"`
}

type ReminderEmailNotificationUsecase interface {
	SendEmail(receivers []string, subject string, body string) error
	PrepareHTMLTemplate(reminder Reminder) (string, error)
	GenerateMessageBuffer(fromAddress mail.Address, to string, cc []string, subject string, body string) bytes.Buffer
}

type ReminderFCMNotificationUsecase interface {
	SendMessage(token string, title string, body string) error
}
