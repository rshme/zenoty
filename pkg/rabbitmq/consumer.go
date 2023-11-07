package rabbitmq

import (
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rafiseptian90/zenoty/utils"
)

type Consumer struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewConsumer() *Consumer {
	conn := NewConnection()
	ch := NewChannel(conn)

	return &Consumer{
		conn: conn,
		ch:   ch,
	}
}

func (c *Consumer) Consume(queueName string, callback func([]byte) error) error {
	q, err := c.ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to declare a queue from %v", queueName)
		return errors.New(errMsg)
	}

	msgs, err := c.ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			err := callback(msg.Body)
			if err != nil {
				msg.Nack(false, true)
				utils.FailOnError("Failed to process the message", err)
			}
			msg.Ack(false)
		}
	}()

	return nil
}

func (c *Consumer) Close() {
	defer c.ch.Close()
	defer c.conn.Close()
}
