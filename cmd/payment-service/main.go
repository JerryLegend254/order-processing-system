package main

import (
	"github.com/JerryLegend254/order-processing-system/internal/logger"
	"github.com/JerryLegend254/order-processing-system/internal/payment"
	"github.com/JerryLegend254/order-processing-system/internal/rabbitmq"
)

func main() {
	logger := logger.New()
	logger.LogInfo("This is the payment service!")

	conn, err := rabbitmq.Connect()
	if err != nil {
		logger.LogError(err, "Failed to connect to RabbitMQ")
		return
	}
	defer conn.Close()

	logger.LogInfo("Connected to RabbitMQ")
	paymentHandler := payment.NewHandler(conn, logger)

	// Listen for order created event
	paymentHandler.ListenForOrder()

	var forever chan struct{}

	<-forever
}
