package application

import "github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/domain"

type TokenGenerator interface {
	Generate(userID domain.UserID) (string, error)
}
