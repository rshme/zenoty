package usecase

import (
	"context"
	"firebase.google.com/go/v4/messaging"
)

type ReminderFCMNotification struct {
	ctx       context.Context
	fcmClient *messaging.Client
}

func NewReminderFCMNotification(ctx context.Context, fcmClient *messaging.Client) ReminderFCMNotification {
	return ReminderFCMNotification{
		ctx:       ctx,
		fcmClient: fcmClient,
	}
}

func (r *ReminderFCMNotification) SendMessage(token string, title string, body string) error {
	_, err := r.fcmClient.Send(r.ctx, &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: token,
	})
	if err != nil {
		return err
	}

	return nil
}
