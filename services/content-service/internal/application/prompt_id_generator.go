package application

import "github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/domain"

type PromptIdGenerator interface {
	Generate() domain.PromptID
}
