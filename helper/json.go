package helper

import (
	"encoding/json"
	"net/http"
)

// WriteToResponseBody encode to json, and as trigger panic handler
func WriteToResponseBody(writer http.ResponseWriter, webResponse interface{}) {
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	PanicIfError(err)
}