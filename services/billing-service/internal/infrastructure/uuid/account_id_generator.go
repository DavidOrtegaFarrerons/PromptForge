package uuid

import (
	"github.com/DavidOrtegaFarrerons/promptforge/services/billing-service/internal/domain"
	"github.com/google/uuid"
)

type AccountIDGenerator struct {
}

func NewAccountIDGenerator() *AccountIDGenerator {
	return &AccountIDGenerator{}
}

func (u *AccountIDGenerator) Generate() domain.AccountID {
	return domain.AccountID(uuid.NewString())
}
