package application

import (
	"context"

	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/domain"
)

type LoginUserInput struct {
	Email    string
	Password string
}

type LoginUserService struct {
	userRepository UserRepository
	passwordHasher PasswordHasher
	tokenGenerator TokenGenerator
}

func NewLoginUserService(userRepository UserRepository, passwordHasher PasswordHasher, tokenGenerator TokenGenerator) *LoginUserService {
	return &LoginUserService{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
	}
}

func (s *LoginUserService) Execute(ctx context.Context, input LoginUserInput) (string, error) {
	email, err := domain.NewEmail(input.Email)
	if err != nil {
		return "", err
	}

	if input.Password == "" {
		return "", domain.ErrEmptyPassword
	}

	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = s.passwordHasher.Compare(user.PasswordHash(), []byte(input.Password))
	if err != nil {
		return "", err
	}

	token, err := s.tokenGenerator.Generate(user.ID(), user.Email())
	if err != nil {
		return "", ErrTokenCouldNotBeGenerated
	}

	return token, nil
}
