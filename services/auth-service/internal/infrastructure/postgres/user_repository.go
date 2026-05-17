package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/application"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/domain"
	"github.com/jackc/pgx/v5/pgconn"
)

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
