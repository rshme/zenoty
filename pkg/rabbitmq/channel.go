package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rafiseptian90/zenoty/utils"
)

func NewChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		utils.FailOnError("Failed to open a channel", err)
	}

	return ch
}
