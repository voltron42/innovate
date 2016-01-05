package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":3020", http.FileServer(http.Dir("./web/"))))
}
