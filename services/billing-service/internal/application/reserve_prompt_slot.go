package application

import (
	"context"

	"github.com/DavidOrtegaFarrerons/promptforge/services/billing-service/internal/domain"
)

type ReservePromptSlotService struct {
	accountRepository AccountRepository
}

func NewReservePromptSlotService(accountRepository AccountRepository) ReservePromptSlotService {
	return ReservePromptSlotService{accountRepository: accountRepository}
}

type ReservePromptSlotInput struct {
	UserID string
}

func (s ReservePromptSlotService) Execute(ctx context.Context, input ReservePromptSlotInput) error {
	if input.UserID == "" {
		return domain.ErrEmptyUserID
	}

	return s.accountRepository.ReservePromptSlot(ctx, domain.UserID(input.UserID))
}
