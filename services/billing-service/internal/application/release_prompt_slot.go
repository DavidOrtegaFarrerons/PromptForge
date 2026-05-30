package application

import (
	"context"

	"github.com/DavidOrtegaFarrerons/promptforge/services/billing-service/internal/domain"
)

type ReleasePromptSlotService struct {
	accountRepository AccountRepository
}

func NewReleasePromptSlotService(accountRepository AccountRepository) ReleasePromptSlotService {
	return ReleasePromptSlotService{accountRepository: accountRepository}
}

type ReleasePromptSlotInput struct {
	UserID string
}

func (s ReleasePromptSlotService) Execute(ctx context.Context, input ReleasePromptSlotInput) error {
	if input.UserID == "" {
		return domain.ErrEmptyUserID
	}

	return s.accountRepository.ReleasePromptSlot(ctx, domain.UserID(input.UserID))
}
