package routes

import (
	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func CreateRouter() *mux.Router {
	return r
}

func createRoutes(router *mux.Router) {
	r.HandleFunc("/", serv1).Methods("GET")
}
