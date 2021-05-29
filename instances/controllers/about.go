package controllers

import (
	"encoding/json"
	"net/http"

	structures "bitbucket.org/burhanmubarok/microservice/structures/https"
)

// About doc
type About struct{}

// Get doc
func (a About) Get(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	data := structures.Response{
		Message: "Microservice [desc]",
	}
	json.NewEncoder(res).Encode(data)
}
