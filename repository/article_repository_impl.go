package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"restfull-api-arcticles/helper"
	"restfull-api-arcticles/model/domain"
)

type AnimeRepositoryImpl struct {
	
}

func NewAnimeAricleRepository() ArticleRepository {
	return &AnimeRepositoryImpl{}
}

func (anime *AnimeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article {

	SQL := "INSERT INTO anime_articles(title, author, content) VALUES (? ,?, ?)"
	result, err := tx.ExecContext(ctx, SQL, article.Title, article.Author, article.Content)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	article.Id = int(id)

	return article

}

func (anime *AnimeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article {
	SQL := "UPDATE anime_articles SET title = (?), author = (?), content = (?) WHERE id = (?)"

	_, err := tx.QueryContext(ctx, SQL, article.Title, article.Author, article.Content, article.Id)
	helper.PanicIfError(err)

	return article

}

func (anime *AnimeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	SQL := "DELETE FROM anime_articles WHERE id = (?)"

	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)

	return nil

}

func (anime *AnimeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Article, error) {
	SQL := "SELECT id, title, author, content, created_at FROM anime_articles WHERE id = ? LIMIT 1"

	rows, err := tx.QueryContext(ctx, SQL, id)
	defer rows.Close()

	helper.PanicIfError(err)

	article := domain.Article{}
	//
	//var idArticle int32
	//var title, author, content string
	//var createdAt time.Time

	if rows.Next() {
		err := rows.Scan(&article.Id, &article.Title, &article.Author, &article.Content, &article.CreatedAt)
		helper.PanicIfError(err)

		return article, nil

	} else {
		return article, errors.New(fmt.Sprintf("data with id %d not found", id))
	}


}

func (anime *AnimeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Article {
	SQL := "SELECT id, title, author, content, created_at FROM anime_articles"

	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()

	helper.PanicIfError(err)

	var articles []domain.Article
	//
	//var idArticle int32
	//var title, author, content string
	//var createdAt time.Time

	for rows.Next() {

		article := domain.Article{}
		err := rows.Scan(&article.Id, &article.Title, &article.Author, &article.Content, &article.CreatedAt)
		helper.PanicIfError(err)

		if (domain.Article{}) != article {
			articles = append(articles, article)
		}

	}

	return articles
}

