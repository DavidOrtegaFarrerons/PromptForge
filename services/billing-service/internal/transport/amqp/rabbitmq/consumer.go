package rabbitmq

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/DavidOrtegaFarrerons/promptforge/services/billing-service/internal/application"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	c                    *amqp.Channel
	createAccountService *application.CreateAccountService
	accountIdGenerator   application.AccountIDGenerator
}

func NewRabbitMQConsumer(c *amqp.Channel, createAccountService *application.CreateAccountService, accountIdGenerator application.AccountIDGenerator) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		c:                    c,
		createAccountService: createAccountService,
		accountIdGenerator:   accountIdGenerator,
	}
}

type UserRegisteredEvent struct {
	UserID     string    `json:"user_id"`
	Email      string    `json:"email"`
	OccurredAt time.Time `json:"occurred_at"`
}

func (c *RabbitMQConsumer) Consume() {
	q, err := c.c.QueueDeclare("billing.user.registered", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = c.c.QueueBind(q.Name, "user.registered", "promptforge.events", false, nil)
	if err != nil {
		panic(err)
	}

	msgs, err := c.c.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	for msg := range msgs {
		var event UserRegisteredEvent
		err := json.NewDecoder(bytes.NewBuffer(msg.Body)).Decode(&event)
		if err != nil {
			log.Printf("Error when decoding: %+v", err.Error())
			msg.Nack(false, false)
			continue
		}

		input := application.CreateAccountInput{
			AccountID: string(c.accountIdGenerator.Generate()),
			UserID:    event.UserID,
			Plan:      "free",
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		_, err = c.createAccountService.Execute(ctx, input)
		if err != nil {
			log.Printf("Error when executing account creation: %+v", err.Error())
			msg.Nack(false, true)
			cancel()
			continue
		}
		msg.Ack(false)
		cancel()
	}

}
