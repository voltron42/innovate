package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(hello)
	err := http.ListenAndServe(":3010",handler)
	log.Fatal(err)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World."))
}
