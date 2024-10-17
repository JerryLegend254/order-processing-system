package order

import (
	"encoding/json"

	"github.com/JerryLegend254/order-processing-system/internal/models"
	"github.com/JerryLegend254/order-processing-system/internal/utils"
	amq "github.com/rabbitmq/amqp091-go"
)

type Handler struct {
	amqConn *amq.Connection
}

func NewHandler(amqConn *amq.Connection) *Handler {
	return &Handler{amqConn: amqConn}
}

func (h *Handler) CreateOrder(order models.Order) error {
	message, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = utils.PublishMessage(h.amqConn, message, "order.created")
	if err != nil {
		return err
	}

	return nil
}
