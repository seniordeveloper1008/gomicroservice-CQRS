package instances

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/burhanmubarok/microservice/instances/controllers"
	"bitbucket.org/burhanmubarok/microservice/structures/https"
	"github.com/gorilla/mux"
)

// Router doc
type Router struct{}

// Handler doc
func (r *Router) Handler() http.Handler {

	router := mux.NewRouter().StrictSlash(true)

	exceptionCtrl := new(controllers.Exception)
	router.NotFoundHandler = http.HandlerFunc(exceptionCtrl.NotFound)
	router.MethodNotAllowedHandler = http.HandlerFunc(exceptionCtrl.NotAllowed)

	router.Path("/").Methods("GET").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		data := structures.Response{
			Message: "Wellcome! trust me, you are a coder",
		}
		json.NewEncoder(res).Encode(data)
	})

	router.Path("/about").Methods("GET").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		data := structures.Response{
			Message: "Microservice [desc]",
		}
		json.NewEncoder(res).Encode(data)
	})

	return router
}
