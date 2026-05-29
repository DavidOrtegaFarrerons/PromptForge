package application

import (
	"context"
	"time"

	"github.com/DavidOrtegaFarrerons/promptforge/services/billing-service/internal/domain"
)

type CreateAccountService struct {
	accountRepository AccountRepository
}

func NewCreateAccountService(accountRepository AccountRepository) *CreateAccountService {
	return &CreateAccountService{accountRepository: accountRepository}
}

type CreateAccountInput struct {
	AccountID string
	UserID    string
	Plan      string
}

func (s *CreateAccountService) Execute(ctx context.Context, input CreateAccountInput) (domain.Account, error) {
	accountID := domain.AccountID(input.AccountID)
	userID := domain.UserID(input.UserID)
	plan := domain.Plan(input.Plan)
	now := time.Now()

	acc, err := domain.NewAccount(accountID, userID, plan, now, now)
	if err != nil {
		return domain.Account{}, err
	}

	return s.accountRepository.Create(ctx, acc)
}
