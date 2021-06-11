package routers

import (
	"net/http"

	"github.com/burhanmubarok/microservice/instances/controllers"
	mid "github.com/burhanmubarok/microservice/instances/middlewares"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Router doc
type Router struct{}

// Handler doc
func (r *Router) Handler() http.Handler {

	routeHandler := mux.NewRouter().StrictSlash(true)

	exceptionCtrl := new(controllers.Exception)
	routeHandler.NotFoundHandler = http.HandlerFunc(exceptionCtrl.NotFound)
	routeHandler.MethodNotAllowedHandler = http.HandlerFunc(exceptionCtrl.NotAllowed)

	routeHandler.Path("/").Methods("GET").HandlerFunc(new(controllers.Home).Get)
	routeHandler.Path("/about").Methods("GET").HandlerFunc(new(controllers.About).Get)

	router := negroni.New()
	router.Use(negroni.HandlerFunc(new(mid.Middleware).Apply))
	router.UseHandler(routeHandler)

	return router
}
