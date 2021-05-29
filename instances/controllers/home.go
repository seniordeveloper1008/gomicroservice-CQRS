package controllers

import (
	"net/http"

	"bitbucket.org/burhanmubarok/microservice/instances/helpers"
	structures "bitbucket.org/burhanmubarok/microservice/structures/https"
)

// Home doc
type Home struct{}

// Get doc
func (h *Home) Get(w http.ResponseWriter, r *http.Request) {
	new(helpers.Response).SendOK(w, structures.Response{Message: "Wellcome to Microservice [Name]"})
}
