package routers

import (
	"net/http"

	"github.com/burhanmubarok/microservice/instances/controllers"
	mid "github.com/burhanmubarok/microservice/instances/middlewares"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Router doc
type Router struct {
	handler *mux.Router
	router  *negroni.Negroni
}

// Init doc
func (r *Router) Init() {
	r.applyHandler()
	r.applyRouter()
	r.applyRoute()
}

// Handler doc
func (r *Router) Handler() http.Handler {
	return r.router
}

func (r *Router) applyHandler() {
	exceptionCtrl := new(controllers.Exception)
	r.handler = mux.NewRouter().StrictSlash(true)
	r.handler.NotFoundHandler = http.HandlerFunc(exceptionCtrl.NotFound)
	r.handler.MethodNotAllowedHandler = http.HandlerFunc(exceptionCtrl.NotAllowed)
}

func (r *Router) applyRouter() {
	r.router = negroni.New()
	r.router.Use(negroni.HandlerFunc(new(mid.Middleware).Apply))
	r.router.UseHandler(r.handler)
}

func (r *Router) applyRoute() {
	r.handler.Path("/").Methods("GET").HandlerFunc(new(controllers.Home).Get)
	r.handler.Path("/about").Methods("GET").HandlerFunc(new(controllers.About).Get)
}
