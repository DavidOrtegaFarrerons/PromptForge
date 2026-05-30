package application

import (
	"context"
	"time"

	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/domain"
)

type CreatePromptInput struct {
	OwnerID string
	Title   string
	Content string
	Tags    []string
}
type CreatePromptService struct {
	promptIdGenerator PromptIdGenerator
	promptRepository  PromptRepository
	billingClient     BillingClient
}

func NewCreatePromptService(promptIdGenerator PromptIdGenerator, promptRepository PromptRepository, billingClient BillingClient) *CreatePromptService {
	return &CreatePromptService{
		promptIdGenerator: promptIdGenerator,
		promptRepository:  promptRepository,
		billingClient:     billingClient,
	}
}

func (s *CreatePromptService) Execute(ctx context.Context, input CreatePromptInput) (domain.Prompt, error) {
	promptID := s.promptIdGenerator.Generate()
	promptTemplate, err := domain.NewPromptTemplate(input.Content)
	if err != nil {
		return domain.Prompt{}, err
	}

	now := time.Now()

	err = s.billingClient.ReservePromptSlot(ctx, input.OwnerID)
	if err != nil {
		return domain.Prompt{}, err
	}
	prompt, err := domain.NewPrompt(promptID, input.OwnerID, input.Title, promptTemplate, input.Tags, now, now)
	if err != nil {
		_ = s.billingClient.ReleasePromptSlot(ctx, input.OwnerID)
		return domain.Prompt{}, err
	}

	prompt, err = s.promptRepository.Create(ctx, prompt)
	if err != nil {
		_ = s.billingClient.ReleasePromptSlot(ctx, input.OwnerID)
		return domain.Prompt{}, err
	}

	return prompt, nil
}
