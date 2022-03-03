package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restfull-api-arcticles/config"
	"restfull-api-arcticles/controller"
	"restfull-api-arcticles/exception"
	"restfull-api-arcticles/middleware"
	"restfull-api-arcticles/repository"
	"restfull-api-arcticles/service"
)

func main() {

	// db := config.GetConnection(config.New())
	db := config.GetConnection(config.New())
	validation := validator.New()

	newArticleAnimeRepository := repository.NewAnimeAricleRepository()
	newArticleAnimeService := service.NewArticleService(newArticleAnimeRepository, db, validation)
	newArticleAnimeController := controller.NewArticleController(newArticleAnimeService)


	router := httprouter.New()

	router.POST("/api/articles", newArticleAnimeController.Create)
	router.GET("/api/articles/:articleId", newArticleAnimeController.FindById)
	router.GET("/api/articles", newArticleAnimeController.FindAll)
	router.DELETE("/api/articles/:articleId", newArticleAnimeController.Delete )
	router.PUT("/api/articles/:articleId", newArticleAnimeController.Update)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:              "localhost:3000",
		Handler:           middleware.NewAuthMiddleware(router),

	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}