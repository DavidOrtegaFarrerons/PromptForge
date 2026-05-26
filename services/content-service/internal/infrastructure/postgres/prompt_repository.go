package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/domain"
	"github.com/lib/pq"
)

type promptRow struct {
	ID        string
	OwnerID   string
	Title     string
	Content   string
	Tags      []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(row promptRow) (domain.Prompt, error) {
	promptTemplate, err := domain.NewPromptTemplate(row.Content)
	if err != nil {
		return domain.Prompt{}, err
	}
	return domain.NewPrompt(
		domain.PromptID(row.ID),
		row.OwnerID,
		row.Title,
		promptTemplate,
		row.Tags,
		row.CreatedAt,
		row.UpdatedAt,
	)
}

type PostgresPromptRepository struct {
	db *sql.DB
}

func NewPostgresPromptRepository(db *sql.DB) *PostgresPromptRepository {
	return &PostgresPromptRepository{db: db}
}

func (r *PostgresPromptRepository) Create(ctx context.Context, prompt domain.Prompt) (domain.Prompt, error) {
	query := `INSERT INTO prompts (
                     id, owner_id, title, content, tags, created_at, updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	args := []any{
		prompt.PromptID(),
		prompt.OwnerID(),
		prompt.Title(),
		prompt.Template().Content(),
		pq.Array(prompt.Tags()),
		prompt.CreatedAt(),
		prompt.UpdatedAt(),
	}

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return domain.Prompt{}, err
	}

	return prompt, nil
}
