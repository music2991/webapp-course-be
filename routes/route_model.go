package routes

import (
	"net/http"
)

type routeModel struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
	Method  string
}
