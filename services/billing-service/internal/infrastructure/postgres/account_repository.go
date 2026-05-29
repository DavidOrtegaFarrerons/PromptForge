package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DavidOrtegaFarrerons/promptforge/services/billing-service/internal/domain"
	"github.com/jackc/pgx/v5/pgconn"
)

type PostgresAccountRepository struct {
	db *sql.DB
}

func NewPostgresAccountRepository(db *sql.DB) *PostgresAccountRepository {
	return &PostgresAccountRepository{db: db}
}

func (r *PostgresAccountRepository) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	query := `INSERT INTO accounts (id, user_id, plan, created_at, updated_at) VALUES
	($1, $2, $3, $4, $5)
`
	_, err := r.db.ExecContext(ctx, query,
		string(account.AccountID()),
		string(account.UserID()),
		string(account.Plan()),
		account.CreatedAt(),
		account.UpdatedAt(),
	)

	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok && pgErr.Code == "23505" {
			return domain.Account{}, domain.ErrAccountAlreadyExists
		}
		return domain.Account{}, err
	}

	return account, nil
}
