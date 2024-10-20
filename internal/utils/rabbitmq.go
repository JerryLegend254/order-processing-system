package utils

import (
	"context"
	"time"

	"github.com/JerryLegend254/order-processing-system/internal/logger"
	amq "github.com/rabbitmq/amqp091-go"
)

type PubSub struct {
	Conn   *amq.Connection
	logger *logger.Logger
}

func NewPubSub(conn *amq.Connection) *PubSub {
	return &PubSub{Conn: conn, logger: logger.New()}
}

func (pbsb *PubSub) PublishMessage(message []byte, queueName string) error {
	ch, err := pbsb.Conn.Channel()
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

	pbsb.logger.LogInfo("Message published to queue: " + queueName)
	return nil

}

func (pbsb *PubSub) ConsumeMessage(queueName string, message string) {
	ch, err := pbsb.Conn.Channel()
	if err != nil {
		pbsb.logger.LogFatal(err, "Failed to open a channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		pbsb.logger.LogFatal(err, "Failed to declare a queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		pbsb.logger.LogFatal(err, "Failed to register a consumer")
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			pbsb.logger.LogInfo("Received a message: " + string(d.Body))
		}
	}()

	pbsb.logger.LogInfo(message)
	pbsb.logger.LogInfo("Waiting for messages. To exit press CTRL+C")
	<-forever
}
