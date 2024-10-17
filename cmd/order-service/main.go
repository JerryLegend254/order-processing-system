package main

import (
	"fmt"

	"github.com/JerryLegend254/order-processing-system/internal/models"
	"github.com/JerryLegend254/order-processing-system/internal/order"
	"github.com/JerryLegend254/order-processing-system/internal/rabbitmq"
)

func main() {
	fmt.Println("This is the order service!")

	conn, err := rabbitmq.Connect()
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
		return
	}
	defer conn.Close()
	fmt.Println("Connected to RabbitMQ")
	orderHandler := order.NewHandler(conn)
	orderHandler.CreateOrder(models.Order{ProductID: 1, Quantity: 2})

}
