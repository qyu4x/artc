package service

import (
	"context"
	"restfull-api-arcticles/model/web"
)

type ArticleService interface {
	Create(ctx context.Context, request web.ArticleCreateRequest) web.ArticleResponse
	Update(ctx context.Context, request web.ArticleUpdateRequest) web.ArticleResponse
	Delete(ctx context.Context, articleId int)
	FindById(ctx context.Context, articleId int) web.ArticleResponse
	FindAll(ctx context.Context) []web.ArticleResponse
}