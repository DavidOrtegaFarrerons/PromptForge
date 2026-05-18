package application

import (
	"context"

	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
	FindByEmail(ctx context.Context, email domain.Email) (domain.User, error)
}
