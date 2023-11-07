package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/rafiseptian90/zenoty/internal/example_app"
	"github.com/rafiseptian90/zenoty/pkg/logger"
	"github.com/rafiseptian90/zenoty/pkg/rabbitmq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Log.Error("Failed to load .env file")
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic : ", r)
		}
	}()

	exampleappConsumer := rabbitmq.NewConsumer()
	defer exampleappConsumer.Close()

	example_app.RunConsumers(exampleappConsumer)

	var forever chan struct{}

	log.Println("Press CTRL + X to exit process manually ...")

	<-forever
}
