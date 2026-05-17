package application

import "github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/domain"

type UserIDGenerator interface {
	Generate() domain.UserID
}
