package model

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Routing interface {
	Path() string
	Handler(w http.ResponseWriter, req *http.Request)
}

func ApplyRoutings(r *mux.Router, routings ...Routing) {
	for _, routing := range routings {
		r.HandleFunc(routing.Path(), routing.Handler)
	}
}
