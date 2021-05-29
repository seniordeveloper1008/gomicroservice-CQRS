package controllers

import (
	"net/http"

	"bitbucket.org/burhanmubarok/microservice/instances/helpers"
	structures "bitbucket.org/burhanmubarok/microservice/structures/https"
)

// About doc
type About struct{}

// Get doc
func (a *About) Get(w http.ResponseWriter, r *http.Request) {
	new(helpers.Response).SendOK(w, structures.Response{Message: "Microservice [desc]"})
}
