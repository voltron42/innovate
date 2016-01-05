package main

import (
	"./app/model"
	"./app/routings"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	h := routings.NewHello("/hello")
	hq := routings.HelloQuery("/hello_again")
	hp := routings.NewHelloPath("hello")
	js := routings.HelloJson("/hello_json")
	r := mux.NewRouter()
	model.ApplyRoutings(r, h, &hq, hp, &js)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	log.Fatal(http.ListenAndServe(":3040", r))
}
