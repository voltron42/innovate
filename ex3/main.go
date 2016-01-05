package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
  r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, "<h2>Hello world.</h2>")
  })
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	log.Fatal(http.ListenAndServe(":3030", r))
}
