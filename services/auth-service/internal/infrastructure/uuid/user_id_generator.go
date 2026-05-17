package uuid

import (
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/domain"
	gouuid "github.com/google/uuid"
)

type UserIDGenerator struct {
}

func NewUserIDGenerator() *UserIDGenerator {
	return &UserIDGenerator{}
}

func (u *UserIDGenerator) Generate() domain.UserID {
	return domain.UserID(gouuid.NewString())
}
