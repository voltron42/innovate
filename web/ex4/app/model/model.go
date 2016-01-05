package model

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Routing interface {
	Path() string
	Method() Method
	Handler(w http.ResponseWriter, req *http.Request)
}

func ApplyRoutings(r *mux.Router, routings ...Routing) {
	for _, routing := range routings {
		r.Path(routing.Path()).Methods(routing.Method().String()).HandlerFunc(routing.Handler)
	}
}
