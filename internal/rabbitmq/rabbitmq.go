package rabbitmq

import (
	"github.com/JerryLegend254/order-processing-system/internal/config"
	"github.com/JerryLegend254/order-processing-system/internal/logger"
	amq "github.com/rabbitmq/amqp091-go"
)

func Connect() (*amq.Connection, error) {
	logger := logger.New()
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.LogPanic(err, "Failed to load config")
		return nil, err
	}
	conn, err := amq.Dial(cfg.RabbitMQURL)
	if err != nil {
		logger.LogPanic(err, "Failed to connect to RabbitMQ")
		return nil, err
	}
	return conn, nil
}
