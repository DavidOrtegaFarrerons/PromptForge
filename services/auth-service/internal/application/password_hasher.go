package application

type PasswordHasher interface {
	Hash(string) ([]byte, error)
}
