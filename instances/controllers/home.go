package controllers

import (
	"net/http"

	"github.com/burhanmubarok/microservice/instances/helpers"
	structures "github.com/burhanmubarok/microservice/structures/https"
)

// Home doc
type Home struct {
	Gety int
}

// Get doc
func (h *Home) Get(w http.ResponseWriter, r *http.Request) {
	responseData := structures.Response{Message: "Wellcome to Microservice <Name>", Data: "Any data should be loaded here"}
	new(helpers.Response).SendOK(w, responseData)
}
