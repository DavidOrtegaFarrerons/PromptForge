package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/application"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/domain"
	"github.com/jackc/pgx/v5/pgconn"
)

type userRow struct {
	ID           string
	Username     string
	Email        string
	PasswordHash []byte
	CreatedAt    time.Time
}

func toDomain(row userRow) (domain.User, error) {
	email, err := domain.NewEmail(row.Email)
	if err != nil {
		return domain.User{}, err
	}

	return domain.NewUser(
		domain.UserID(row.ID),
		row.Username,
		email,
		row.PasswordHash,
	)
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	query := `INSERT INTO users (id, username, email, password_hash)
	VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, query,
		user.ID(),
		user.Username(),
		user.Email().Value(),
		user.PasswordHash(),
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.User{}, application.ErrDuplicateEmail
		}

		return domain.User{}, err
	}

	return user, nil
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email domain.Email) (domain.User, error) {
	query := `SELECT id, username, email, password_hash FROM users WHERE email = $1`

	var user userRow
	err := r.db.QueryRowContext(ctx, query, email.Value()).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return toDomain(user)
}
