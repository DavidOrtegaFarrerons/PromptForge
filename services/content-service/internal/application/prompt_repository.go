package application

import (
	"context"

	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/domain"
)

type PromptRepository interface {
	Create(ctx context.Context, prompt domain.Prompt) (domain.Prompt, error)
}
