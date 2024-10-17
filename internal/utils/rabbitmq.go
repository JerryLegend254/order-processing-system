package utils

import (
	"context"
	"log"
	"time"

	amq "github.com/rabbitmq/amqp091-go"
)

func PublishMessage(conn *amq.Connection, message []byte, queueName string) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amq.Publishing{
		ContentType: "application/json",
		Body:        message,
	})

	if err != nil {
		return err
	}

	log.Printf("Message published to queue: %s", queueName)
	return nil

}
