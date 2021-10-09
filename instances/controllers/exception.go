package controllers

import (
	"net/http"

	"github.com/burubur/microservice/instances/helpers"
	structures "github.com/burubur/microservice/structures/https"
)

// Exception doc
type Exception struct{}

// NotFound 404
func (e *Exception) NotFound(w http.ResponseWriter, r *http.Request) {
	new(helpers.Response).SendError(w, http.StatusNotFound, structures.Response{Message: "URL not found"})
}

// NotAllowed 405
func (e *Exception) NotAllowed(w http.ResponseWriter, r *http.Request) {
	new(helpers.Response).SendError(w, http.StatusNotFound, structures.Response{Message: "Method not allowed for that end point"})
}

// InternalServerError 500
func (e *Exception) InternalServerError(w http.ResponseWriter, r *http.Request) {
	new(helpers.Response).SendError(w, http.StatusInternalServerError, structures.Response{Message: "Internal server error"})
}

// BadGateway 502
func (e *Exception) BadGateway(w http.ResponseWriter, r *http.Request) {
	new(helpers.Response).SendError(w, http.StatusBadGateway, structures.Response{Message: "Bad gateway"})
}
