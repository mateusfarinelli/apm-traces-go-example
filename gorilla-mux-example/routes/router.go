package routes

import (
	"apm-trace-with-gorilla-mux-example/dependencies"
	"apm-trace-with-gorilla-mux-example/errs"
	"apm-trace-with-gorilla-mux-example/handles"
	"net/http"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

type Route struct {
	URI string
	Method string
	Function func(http.ResponseWriter, *http.Request)
}

type Router struct {
	Router *muxtrace.Router
}


func NewRouter() *muxtrace.Router {
	router := Router{Router: muxtrace.NewRouter(muxtrace.WithServiceName("apm-trace-example-with-gorilla-mux"))}

	router.LoadRoutes()

	return router.Router
}

func (r *Router) LoadRoutes(){
	r.ConfigRoutes(GetProductRoutes(dependencies.GetProductController()))
	r.Router.NotFoundHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handles.Error(w, errs.NotFound("Resource Not Found"))
		},
	)
}

func (r *Router) ConfigRoutes(rt []Route) {
	for _, route := range rt {
		r.Router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
}