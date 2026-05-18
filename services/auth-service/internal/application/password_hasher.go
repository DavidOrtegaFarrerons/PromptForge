package application

type PasswordHasher interface {
	Hash(password string) ([]byte, error)
	Compare(hash, password []byte) error
}
