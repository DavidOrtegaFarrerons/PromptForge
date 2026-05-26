package uuid

import (
	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/domain"
	"github.com/google/uuid"
)

type PromptIDGenerator struct {
}

func NewPromptIdGenerator() *PromptIDGenerator {
	return &PromptIDGenerator{}
}

func (g *PromptIDGenerator) Generate() domain.PromptID {
	return domain.PromptID(uuid.NewString())
}
