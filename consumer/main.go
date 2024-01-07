package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

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

	emailConsumer, err := ch.ConsumeWithContext(
		ctx,
		"email",          // queue name
		"email-consumer", // consumer name
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	for msg := range emailConsumer {
		log.Printf("Routing key: %s", msg.RoutingKey)
		log.Printf("Received message: %s", msg.Body)
		// err = msg.Ack(false)
		// if err != nil {
		// 	log.Fatalf("Failed to acknowledge message: %s", err)
		// }
	}
}
