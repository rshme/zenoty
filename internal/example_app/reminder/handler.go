package reminder

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rafiseptian90/zenoty/internal/example_app/config"
	"github.com/rafiseptian90/zenoty/internal/example_app/domain"
	"github.com/rafiseptian90/zenoty/internal/example_app/reminder/usecase"
	"github.com/rafiseptian90/zenoty/pkg/logger"
)

type ReminderNotificationHandler struct {
	mailConfig config.SMTPMailConfig
	fcmConfig  config.FCMConfig
}

func NewReminderNotificationHandler(mailConfig config.SMTPMailConfig, fcmConfig config.FCMConfig) *ReminderNotificationHandler {
	return &ReminderNotificationHandler{
		mailConfig: mailConfig,
		fcmConfig:  fcmConfig,
	}
}

func (h *ReminderNotificationHandler) EmailHandler(msg []byte) error {
	var reminder domain.Reminder
	if err := json.Unmarshal(msg, &reminder); err != nil {
		logger.Log.Errorf("Failed to unmarshal JSON : %v", err)
	}

	reminderEmailUsecase := usecase.NewReminderEmailNotification(h.mailConfig)

	body, err := reminderEmailUsecase.PrepareHTMLTemplate(reminder)
	if err != nil {
		return err
	}

	if err := reminderEmailUsecase.SendEmail(strings.Split(reminder.Email, "|"), "Example App Task Reminder", body); err != nil {
		return err
	}

	logger.Log.Info("Successfully sent reminder email notification...")

	return nil
}

func (h *ReminderNotificationHandler) FCMHandler(msg []byte) error {
	ctx := context.Background()

	var reminder domain.Reminder
	if err := json.Unmarshal(msg, &reminder); err != nil {
		logger.Log.Errorf("Failed to unmarshal JSON : %v", err)
	}

	client, err := h.fcmConfig.NewFCMClient(ctx)
	if err != nil {
		logger.Log.Errorf("Failed to create a FCM Client : %v\n", err)
		return err
	}

	msgBody := fmt.Sprintf("Hi %s, I want to inform you that you have %d tasks remain today.", reminder.Name, len(reminder.Tasks))
	reminderFCMUsecase := usecase.NewReminderFCMNotification(ctx, client)

	if err := reminderFCMUsecase.SendMessage(reminder.Token, "Example App Task Reminder", msgBody); err != nil {
		logger.Log.Errorf("Failed to sent a reminder FCM message : %v", err)
		return err
	}

	logger.Log.Info("Successfully to sent reminder FCM message")

	return nil
}
