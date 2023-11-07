package rabbitmq

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rafiseptian90/zenoty/utils"
)

func NewConnection() *amqp.Connection {
	url := fmt.Sprintf("amqp://%v:%v@%v:%v", os.Getenv("RABBITMQ_USERNAME"), os.Getenv("RABBITMQ_PASSWORD"), os.Getenv("RABBITMQ_HOST"), os.Getenv("RABBITMQ_PORT"))

	conn, err := amqp.Dial(url)
	if err != nil {
		utils.FailOnError("Failed to connect to RabbitMQ", err)
	}

	return conn
}
