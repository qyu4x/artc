package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restfull-api-arcticles/helper"
	"restfull-api-arcticles/model/web"
	"restfull-api-arcticles/service"
	"strconv"
)

type ArticleControllerImpl struct {
	Service service.ArticleService
}

func NewArticleController(serviceArticle service.ArticleService) ArticleController{
	return &ArticleControllerImpl{
		Service: serviceArticle,
	}
}

func (controller ArticleControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)

	controllerRequestBody := web.ArticleCreateRequest{}

	err := decoder.Decode(&controllerRequestBody)
	helper.PanicIfError(err)

	articleResponse := controller.Service.Create(request.Context(), controllerRequestBody)

	writer.Header().Add("Content-Type", "application/json")
	webResponseWriter := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}
	// encode json
	helper.WriteToResponseBody(writer, webResponseWriter)

}

func (controller ArticleControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)

	controllerRequestBody := web.ArticleUpdateRequest{}

	err := decoder.Decode(&controllerRequestBody)
	helper.PanicIfError(err)

	paramsId := params.ByName("articleId")
	id, err := strconv.Atoi(paramsId)
	helper.PanicIfError(err)

	controllerRequestBody.Id = id

	articleResponse := controller.Service.Update(request.Context(), controllerRequestBody)

	writer.Header().Add("Content-Type", "application/json")
	webResponseWriter := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}

	// encode json
	helper.WriteToResponseBody(writer, webResponseWriter)
}

func (controller ArticleControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {


	controllerRequestBody := web.ArticleUpdateRequest{}

	paramsId := params.ByName("articleId")
	id, err := strconv.Atoi(paramsId)
	helper.PanicIfError(err)

	controllerRequestBody.Id = id

	controller.Service.Delete(request.Context(), controllerRequestBody.Id)

	writer.Header().Add("Content-Type", "application/json")
	webResponseWriter := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	// encode json
	helper.WriteToResponseBody(writer, webResponseWriter)
}

func (controller ArticleControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	controllerRequestBody := web.ArticleUpdateRequest{}

	paramsId := params.ByName("articleId")
	id, err := strconv.Atoi(paramsId)
	helper.PanicIfError(err)

	controllerRequestBody.Id = id

	articleResponse := controller.Service.FindById(request.Context(), controllerRequestBody.Id)

	writer.Header().Add("Content-Type", "application/json")
	webResponseWriter := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}

	// encode json
	helper.WriteToResponseBody(writer, webResponseWriter)
}

func (controller ArticleControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	articleResponse := controller.Service.FindAll(request.Context())

	writer.Header().Add("Content-Type", "application/json")
	webResponseWriter := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}

	// encode json
	helper.WriteToResponseBody(writer, webResponseWriter)
}
