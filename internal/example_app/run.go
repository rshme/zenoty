package example_app

import (
	"github.com/rafiseptian90/zenoty/internal/example_app/config"
	"github.com/rafiseptian90/zenoty/internal/example_app/reminder"
	"github.com/rafiseptian90/zenoty/pkg/logger"
	"github.com/rafiseptian90/zenoty/pkg/rabbitmq"
)

func RunConsumers(consumer *rabbitmq.Consumer) {

	mailConfig := config.NewSMTPConfiguration()
	fcmConfig := config.NewFCMConfig("config/fcm_config.json")

	reminderHandler := reminder.NewReminderNotificationHandler(mailConfig, fcmConfig)

	if err := consumer.Consume("exampleapp.reminder.email_queue", reminderHandler.EmailHandler); err != nil {
		logger.Log.Error(err.Error())
	}

	if err := consumer.Consume("exampleapp.reminder.fcm_queue", reminderHandler.FCMHandler); err != nil {
		logger.Log.Error(err.Error())
	}
}
