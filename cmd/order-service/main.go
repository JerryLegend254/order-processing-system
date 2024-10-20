package main

import (
	"github.com/JerryLegend254/order-processing-system/internal/logger"
	"github.com/JerryLegend254/order-processing-system/internal/models"
	"github.com/JerryLegend254/order-processing-system/internal/order"
	"github.com/JerryLegend254/order-processing-system/internal/rabbitmq"
)

func main() {
	logger := logger.New()
	logger.LogInfo("This is the order service!")

	conn, err := rabbitmq.Connect()
	if err != nil {
		logger.LogError(err, "Failed to connect to RabbitMQ")
		return
	}
	defer conn.Close()

	logger.LogInfo("Connected to RabbitMQ")

	orderHandler := order.NewHandler(conn, logger)
	orderHandler.CreateOrder(models.Order{ProductID: 1, Quantity: 2})

}
