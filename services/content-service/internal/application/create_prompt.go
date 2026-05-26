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
}

func NewCreatePromptService(promptIdGenerator PromptIdGenerator, promptRepository PromptRepository) *CreatePromptService {
	return &CreatePromptService{
		promptIdGenerator: promptIdGenerator,
		promptRepository:  promptRepository,
	}
}

func (s *CreatePromptService) Execute(ctx context.Context, input CreatePromptInput) (domain.Prompt, error) {
	promptID := s.promptIdGenerator.Generate()
	promptTemplate, err := domain.NewPromptTemplate(input.Content)
	if err != nil {
		return domain.Prompt{}, err
	}

	now := time.Now()
	prompt, err := domain.NewPrompt(promptID, input.OwnerID, input.Title, promptTemplate, input.Tags, now, now)
	if err != nil {
		return domain.Prompt{}, err
	}

	return s.promptRepository.Create(ctx, prompt)
}
