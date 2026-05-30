package application

import (
	"context"

	"github.com/DavidOrtegaFarrerons/promptforge/services/billing-service/internal/domain"
)

type AccountRepository interface {
	Create(ctx context.Context, account domain.Account) (domain.Account, error)
	ReservePromptSlot(ctx context.Context, userID domain.UserID) error
	ReleasePromptSlot(ctx context.Context, userID domain.UserID) error
}
