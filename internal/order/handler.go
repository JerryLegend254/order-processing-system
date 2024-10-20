package order

import (
	"encoding/json"

	"github.com/JerryLegend254/order-processing-system/internal/logger"
	"github.com/JerryLegend254/order-processing-system/internal/models"
	"github.com/JerryLegend254/order-processing-system/internal/utils"
	amq "github.com/rabbitmq/amqp091-go"
)

type Handler struct {
	amqConn *amq.Connection
	logger  *logger.Logger
}

func NewHandler(amqConn *amq.Connection, logger *logger.Logger) *Handler {
	return &Handler{amqConn: amqConn, logger: logger}
}

func (h *Handler) CreateOrder(order models.Order) error {
	message, err := json.Marshal(order)
	if err != nil {
		return err
	}

	pbsb := utils.NewPubSub(h.amqConn)
	err = pbsb.PublishMessage(message, "order.created")
	if err != nil {
		return err
	}

	return nil
}
