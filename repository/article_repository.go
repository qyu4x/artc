package repository

import (
	"context"
	"database/sql"
	"restfull-api-arcticles/model/domain"
)

type ArticleRepository interface {
	Save(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article
	Update(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Article, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Article
}
