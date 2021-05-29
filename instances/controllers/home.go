package controllers

import (
	"encoding/json"
	"net/http"

	structures "bitbucket.org/burhanmubarok/microservice/structures/https"
)

// Home doc
type Home struct{}

// Get doc
func (h Home) Get(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	data := structures.Response{
		Message: "Wellcome! trust me, you are a coder",
	}
	json.NewEncoder(res).Encode(data)
}
