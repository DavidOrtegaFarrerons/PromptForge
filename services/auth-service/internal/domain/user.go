package domain

type UserID string

type User struct {
	id           UserID
	username     string
	email        Email
	passwordHash []byte
}

func NewUser(id UserID, username string, email Email, passwordHash []byte) (User, error) {
	if id == "" {
		return User{}, ErrInvalidUser
	}

	if username == "" {
		return User{}, ErrInvalidUser
	}

	if email.IsZero() {
		return User{}, ErrInvalidEmail
	}

	if len(passwordHash) == 0 {
		return User{}, ErrInvalidUser
	}

	return User{
		id:           id,
		username:     username,
		email:        email,
		passwordHash: passwordHash,
	}, nil
}

func (u User) ID() UserID {
	return u.id
}

func (u User) Username() string {
	return u.username
}

func (u User) Email() Email {
	return u.email
}

func (u User) PasswordHash() []byte {
	return u.passwordHash
}
