package controllers

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/burhanmubarok/microservice/structures/https"
)

// Exception doc
type Exception struct{}

// NotFound 404
func (e *Exception) NotFound(w http.ResponseWriter, r *http.Request) {
	e.response(w, http.StatusNotFound, "URL not found", nil, nil)
}

// NotAllowed 405
func (e *Exception) NotAllowed(w http.ResponseWriter, r *http.Request) {
	e.response(w, http.StatusMethodNotAllowed, "Method not allowed for that end point", nil, nil)
}

// InternalServerError 500
func (e *Exception) InternalServerError(w http.ResponseWriter, r *http.Request) {
	e.response(w, http.StatusInternalServerError, "Internal server error", nil, nil)
}

// BadGateway 502
func (e *Exception) BadGateway(w http.ResponseWriter, r *http.Request) {
	e.response(w, http.StatusBadGateway, "Bad gateway", nil, nil)
}

func (e *Exception) response(w http.ResponseWriter, httpStatus int, message interface{}, data interface{}, error interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	responseData := structures.Response{
		Message: message,
		Data:    data,
		Errors:  error,
	}
	json.NewEncoder(w).Encode(responseData)
}
