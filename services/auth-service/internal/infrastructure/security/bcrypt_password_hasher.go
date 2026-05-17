package security

import "golang.org/x/crypto/bcrypt"

type BcryptPasswordHasher struct {
	cost int
}

func NewBcryptPasswordHasher(cost int) *BcryptPasswordHasher {
	return &BcryptPasswordHasher{
		cost: cost,
	}
}

func (h *BcryptPasswordHasher) Hash(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
