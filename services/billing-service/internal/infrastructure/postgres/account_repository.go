package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DavidOrtegaFarrerons/promptforge/services/billing-service/internal/domain"
	"github.com/jackc/pgx/v5/pgconn"
)

type accountRow struct {
	ID          string
	UserID      string
	Plan        string
	PromptCount int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func toDomain(row accountRow) (domain.Account, error) {
	return domain.NewAccount(
		domain.AccountID(row.ID),
		domain.UserID(row.UserID),
		domain.Plan(row.Plan),
		row.PromptCount,
		row.CreatedAt,
		row.UpdatedAt,
	)
}

type PostgresAccountRepository struct {
	db *sql.DB
}

func NewPostgresAccountRepository(db *sql.DB) *PostgresAccountRepository {
	return &PostgresAccountRepository{db: db}
}

func (r *PostgresAccountRepository) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	query := `INSERT INTO accounts (id, user_id, plan, prompt_count, created_at, updated_at) VALUES
	($1, $2, $3, $4, $5, $6)
`
	_, err := r.db.ExecContext(ctx, query,
		string(account.AccountID()),
		string(account.UserID()),
		string(account.Plan()),
		account.PromptCount(),
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

func (r *PostgresAccountRepository) ReservePromptSlot(ctx context.Context, userID domain.UserID) error {
	lockQuery := `SELECT id, user_id, plan, prompt_count, created_at, updated_at FROM accounts WHERE user_id = $1 FOR UPDATE`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var accRow accountRow
	err = tx.QueryRowContext(ctx, lockQuery, string(userID)).Scan(
		&accRow.ID,
		&accRow.UserID,
		&accRow.Plan,
		&accRow.PromptCount,
		&accRow.CreatedAt,
		&accRow.UpdatedAt,
	)
	if err != nil {
		return err
	}

	acc, err := toDomain(accRow)
	if err != nil {
		return err
	}

	if !acc.CanCreatePrompt() {
		return domain.ErrPromptLimitReached
	}

	updateQuery := `UPDATE accounts SET prompt_count = prompt_count + 1 WHERE user_id = $1`

	_, err = tx.ExecContext(ctx, updateQuery, string(userID))
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresAccountRepository) ReleasePromptSlot(ctx context.Context, userID domain.UserID) error {
	query := `UPDATE accounts SET prompt_count = prompt_count - 1 WHERE user_id = $1`

	_, err := r.db.ExecContext(ctx, query, string(userID))
	return err
}
