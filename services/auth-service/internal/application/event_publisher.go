package application

import (
	"context"

	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/domain"
)

type EventPublisher interface {
	PublishUserRegistered(ctx context.Context, event domain.UserRegisteredEvent) error
}
