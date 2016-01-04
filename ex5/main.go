package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/calendar/{year}/{month}", func(w http.ResponseWriter, req *http.Request) {

	})
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	log.Fatal(http.ListenAndServe(":3030", r))
}
