package helpers

import (
	"encoding/json"
	"net/http"

	structures "bitbucket.org/burhanmubarok/microservice/structures/https"
)

// Response doc
type Response struct{}

// SendOK doc
func (r *Response) SendOK(w http.ResponseWriter, data structures.Response) {
	r.send(w, http.StatusOK, data)
}

// SendAccepted doc
func (r *Response) SendAccepted(w http.ResponseWriter, data structures.Response) {
	r.send(w, http.StatusAccepted, data)
}

func (r *Response) send(w http.ResponseWriter, httpStatus int, data structures.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	responseData := structures.Response{
		Message: data.Message,
		Data:    data.Data,
		Errors:  data.Errors,
	}
	json.NewEncoder(w).Encode(responseData)
}
