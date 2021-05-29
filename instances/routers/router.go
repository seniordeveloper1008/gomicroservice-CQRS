package routers

import (
	"net/http"

	"bitbucket.org/burhanmubarok/microservice/instances/controllers"
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

	router.Path("/").Methods("GET").HandlerFunc(new(controllers.Home).Get)
	router.Path("/about").Methods("GET").HandlerFunc(new(controllers.About).Get)

	return router
}
