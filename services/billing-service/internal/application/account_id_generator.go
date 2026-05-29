package application

import "github.com/DavidOrtegaFarrerons/promptforge/services/billing-service/internal/domain"

type AccountIDGenerator interface {
	Generate() domain.AccountID
}
