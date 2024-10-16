package rabbitmq

import (
	"log"

	"github.com/JerryLegend254/order-processing-system/internal/config"
	amq "github.com/rabbitmq/amqp091-go"
)

func Connect() (*amq.Connection, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panicf("Failed to load config: %v", err)
		return nil, err
	}
	conn, err := amq.Dial(cfg.RabbitMQURL)
	if err != nil {
		log.Panicf("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}
	return conn, nil
}
