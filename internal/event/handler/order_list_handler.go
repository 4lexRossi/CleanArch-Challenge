package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/4lexRossi/CleanArch-Challenge/pkg/events"
	"github.com/streadway/amqp"
)

type OrderListHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderListHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) ListOrders(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order List: %v", event.GetResponse())
	jsonOutput, _ := json.Marshal(event.GetResponse())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
