package payment

import (
	"github.com/JerryLegend254/order-processing-system/internal/logger"
	"github.com/JerryLegend254/order-processing-system/internal/utils"
	amq "github.com/rabbitmq/amqp091-go"
)

type Handler struct {
	amqpConn *amq.Connection
	logger   *logger.Logger
}

func NewHandler(conn *amq.Connection, logger *logger.Logger) *Handler {
	return &Handler{amqpConn: conn, logger: logger}
}

func (h *Handler) ListenForOrder() {
	pbsb := utils.NewPubSub(h.amqpConn)
	pbsb.ConsumeMessage("order.created", "Listening for order created event")
}
