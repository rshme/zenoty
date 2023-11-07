package config

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type FCMConfig struct {
	fcmConfigPath string
}

func NewFCMConfig(fcmConfigPath string) FCMConfig {
	return FCMConfig{fcmConfigPath: fcmConfigPath}
}

func (c *FCMConfig) NewFCMClient(ctx context.Context) (*messaging.Client, error) {
	opt := option.WithCredentialsFile(c.fcmConfigPath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
