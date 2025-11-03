package article

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type Article struct {
	Id          int       `db:"id"`
	CreatedDate time.Time `db:"created_date"`
	Title       string    `db:"title"`
}

type CrudStorage interface {
	Create(context.Context, *Article) error
	Get(context.Context, *Article) error
}

type PostgresRepository struct {
	db *sqlx.DB
}

func NewArticleRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(ctx context.Context, a *Article) error {
	err := r.db.QueryRowContext(ctx, queryInsert, a.CreatedDate, a.Title).Scan(&a.Id)
	return err
}

const queryInsert = `
INSERT INTO article (
	created_date,
	title
) VALUES ($1, $2)
RETURNING id
`

func (r *PostgresRepository) Get(ctx context.Context, a *Article) error {
	err := r.db.GetContext(ctx, a, querySelectById, a.Id)
	if err != nil {
		return err
	}
	return nil
}

const querySelectById = `
SELECT
    id,
	title,
	created_date
FROM article
WHERE id = $1
`
