package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// send sends a message to the queue
func main() {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ server: %s", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create a channel: %s", err)
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for i := 0; i < 10; i++ {
		msg := amqp.Publishing{
			Body: []byte("Hello World!"),
			Headers: amqp.Table{
				"sample": "value",
			},
		}

		err = ch.PublishWithContext(
			ctx,
			"notification", // exchange
			"email",        // routing key
			false,          // mandatory
			false,          // immediate
			msg,
		)

		if err != nil {
			log.Fatalf("Failed to publish a message: %s", err)
		}
	}

	log.Println("Successfully published a message")
}
