package exception

import (
	"github.com/go-playground/validator"
	"net/http"
	"restfull-api-arcticles/helper"
	"restfull-api-arcticles/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if newValidationError(writer, request, err) {
		return
	}

	if newNotFoundError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func newValidationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	errors, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		responseCode := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST NOT FOUND",
			Data:   errors.Error(),
		}
		helper.WriteToResponseBody(writer, responseCode)

		return true
	} else {
		return false
	}
}

func newNotFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	notFoundError, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		responseCode := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "BAD REQUEST NOT FOUND",
			Data:   notFoundError.Error,
		}
		helper.WriteToResponseBody(writer, responseCode)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	responseCode := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, responseCode)
}

