package application

import (
	"context"

	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/domain"
)

type RegisterUserInput struct {
	Username string
	Email    string
	Password string
}

type RegisterUserService struct {
	userRepository  UserRepository
	passwordHasher  PasswordHasher
	userIdGenerator UserIDGenerator
}

func NewRegisterUserService(userRepository UserRepository, passwordHasher PasswordHasher, userIDGenerator UserIDGenerator) *RegisterUserService {
	return &RegisterUserService{
		userRepository:  userRepository,
		passwordHasher:  passwordHasher,
		userIdGenerator: userIDGenerator,
	}
}

func (s *RegisterUserService) Execute(ctx context.Context, input RegisterUserInput) (domain.User, error) {
	email, err := domain.NewEmail(input.Email)
	if err != nil {
		return domain.User{}, err
	}

	passwordHash, err := s.passwordHasher.Hash(input.Password)
	if err != nil {
		return domain.User{}, err
	}

	user, err := domain.NewUser(
		s.userIdGenerator.Generate(),
		input.Username,
		email,
		passwordHash,
	)

	if err != nil {
		return domain.User{}, err
	}

	return s.userRepository.Create(ctx, user)
}
