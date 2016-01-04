package main

import (
	"./app/model"
	"./app/routings"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func main() {
	r := mux.NewRouter()
	model.ApplyRoutings(
		r,
		routings.Hello{"/hello"},
		&routings.HelloQuery(hello_again),
		routings.NewHelloPath("hello"),
	)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	log.Fatal(http.ListenAndServe(":3030", r))
}
