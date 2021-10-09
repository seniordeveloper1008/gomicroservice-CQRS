package controllers

import (
	"net/http"

	"github.com/burubur/microservice/instances/helpers"
	structures "github.com/burubur/microservice/structures/https"
)

// About doc
type About struct{}

// Get doc
func (a *About) Get(w http.ResponseWriter, r *http.Request) {
	new(helpers.Response).SendOK(w, structures.Response{Data: struct {
		Env     string `json:"env"`
		Version string `json:"version"`
	}{
		Env:     "development",
		Version: "0.0.1",
	}})
}
