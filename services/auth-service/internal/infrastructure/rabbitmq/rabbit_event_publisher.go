package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitEventPublisher struct {
	mq *amqp.Channel
}

func NewRabbitEventPublisher(mq *amqp.Channel) *RabbitEventPublisher {
	return &RabbitEventPublisher{mq: mq}
}

func (p *RabbitEventPublisher) PublishUserRegistered(ctx context.Context, event domain.UserRegisteredEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return p.mq.PublishWithContext(ctx, "promptforge.events", "user.registered", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}
