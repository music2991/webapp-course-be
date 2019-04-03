package routes

import (
	"github.com/gorilla/mux"
)

var r *mux.Router

func CreateRouter() *mux.Router {
	r = mux.NewRouter()
	var routes []routeModel
	routes = append(routes, getProductRoutes()...)

	for _, route := range routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
	return r
}
