package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"restfull-api-arcticles/exception"
	"restfull-api-arcticles/helper"
	"restfull-api-arcticles/model/domain"
	"restfull-api-arcticles/model/web"
	"restfull-api-arcticles/repository"
	"time"
)

type ArticleServiceImpl struct {
	Repository repository.ArticleRepository
	DB *sql.DB
	Validation *validator.Validate
}

func NewArticleService(repositoryArticleAnime repository.ArticleRepository, db *sql.DB, validation *validator.Validate) ArticleService{
	return &ArticleServiceImpl{
		Repository: repositoryArticleAnime,
		DB:         db,
		Validation: validation,
	}
}
func (service ArticleServiceImpl) Create(ctx context.Context, request web.ArticleCreateRequest) web.ArticleResponse {
	// validation
	err := service.Validation.Struct(request)
	helper.PanicIfError(err)

	tx, err :=  service.DB.Begin()
	helper.PanicIfError(err)

	// handle transaction
	defer helper.ServiceArticleTransaction(tx)

	articleRepositoryConf := domain.Article{
		Title: request.Title,
		Author: request.Author,
		Content: request.Content,
		CreatedAt: time.Now(),
	}

	articleRepositoryResponse := service.Repository.Save(ctx, tx, articleRepositoryConf)

	articleServiceResponse := web.ArticleResponse{
		Id: articleRepositoryResponse.Id,
		Title: articleRepositoryResponse.Title,
		Author: articleRepositoryResponse.Author,
		Content: articleRepositoryResponse.Content,
		CreatedAt: articleRepositoryResponse.CreatedAt,
	}

	return articleServiceResponse


}

func (service ArticleServiceImpl) Update(ctx context.Context, request web.ArticleUpdateRequest) web.ArticleResponse {
	// validation
	err := service.Validation.Struct(request)

	tx, err :=  service.DB.Begin()
	helper.PanicIfError(err)

	// handle transaction
	defer helper.ServiceArticleTransaction(tx)

	// check data availability
	response, err := service.Repository.FindById(ctx, tx, request.Id)
	if err != nil {
		// panic handler
		panic(exception.NewNotFoundError(err.Error()))
	}

	response.Title = request.Title
	response.Author = request.Author
	response.Content = request.Content

	articleRepositoryResponse := service.Repository.Update(ctx, tx, response)

	articleServiceResponse := web.ArticleResponse{
		Id: articleRepositoryResponse.Id,
		Title: articleRepositoryResponse.Title,
		Author: articleRepositoryResponse.Author,
		Content: articleRepositoryResponse.Content,
		CreatedAt: articleRepositoryResponse.CreatedAt,
	}

	return articleServiceResponse
}

func (service ArticleServiceImpl) Delete(ctx context.Context, articleId int) {
	context := context.Background()
	tx, err :=  service.DB.Begin()
	helper.PanicIfError(err)

	// handle transaction
	defer helper.ServiceArticleTransaction(tx)

	// check data availability
	response, err := service.Repository.FindById(ctx, tx, articleId)
	if err != nil {
		// panic handler
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = service.Repository.Delete(context, tx, response.Id)
	helper.PanicIfError(err)
}

func (service ArticleServiceImpl) FindById(ctx context.Context, articleId int) web.ArticleResponse {
	tx, err :=  service.DB.Begin()
	helper.PanicIfError(err)

	// handle transaction
	defer helper.ServiceArticleTransaction(tx)

	articleRepositoryResponse, err := service.Repository.FindById(ctx, tx, articleId)
	if err != nil {
		// panic handler
		panic(exception.NewNotFoundError(err.Error()))
	}

	articleServiceResponse := web.ArticleResponse{
		Id: articleRepositoryResponse.Id,
		Title: articleRepositoryResponse.Title,
		Author: articleRepositoryResponse.Author,
		Content: articleRepositoryResponse.Content,
		CreatedAt: articleRepositoryResponse.CreatedAt,
	}

	return articleServiceResponse
}

func (service ArticleServiceImpl) FindAll(ctx context.Context) []web.ArticleResponse {
	tx, err :=  service.DB.Begin()
	helper.PanicIfError(err)

	// handle transaction
	defer helper.ServiceArticleTransaction(tx)

	articleServiceResponses :=  []web.ArticleResponse{}

	articleRepositoryResponses := service.Repository.FindAll(ctx, tx)
	helper.PanicIfError(err)

	for _, articleResponse := range articleRepositoryResponses {
		articleServiceResponse :=  web.ArticleResponse{
			Id: articleResponse.Id,
			Title: articleResponse.Title,
			Author: articleResponse.Author,
			Content: articleResponse.Content,
			CreatedAt: articleResponse.CreatedAt,
		}

		articleServiceResponses = append(articleServiceResponses, articleServiceResponse)

	}

	return articleServiceResponses
}
