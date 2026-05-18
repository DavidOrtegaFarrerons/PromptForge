package security

import "golang.org/x/crypto/bcrypt"

type BcryptPasswordHasher struct {
	cost int
}

func NewBcryptPasswordHasher() *BcryptPasswordHasher {
	return &BcryptPasswordHasher{
		cost: bcrypt.DefaultCost,
	}
}

func (h *BcryptPasswordHasher) Hash(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func (h *BcryptPasswordHasher) Compare(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
