package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"restfull-api-arcticles/helper"
	"restfull-api-arcticles/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}


func (middeware *AuthMiddleware)ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	os.Setenv(
		"API_KEY",
		"SHIYZU",
	)
	apiKeyMiddleware := os.Getenv("API_KEY")

	if apiKeyMiddleware != request.Header.Get("x-api-key") {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		responseCode := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTORIZED",
			Data:    errors.New(fmt.Sprintf("Please enter a correect key in the header")),
		}
		helper.WriteToResponseBody(writer, responseCode)
	} else {
		middeware.Handler.ServeHTTP(writer, request)
	}
}

