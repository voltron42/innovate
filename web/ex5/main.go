package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
  "./factors"
  "math/rand"
  "encoding/json"
)

func main() {
	r := mux.NewRouter()
  r.HandleFunc("/statdata", func(w http.ResponseWriter, req *http.Request) {
    value := rand.Intn(98) + 2
    f := factors.Factor(value)
    output := map[string]interface{}{
      "selected":value,
      "factors":f,
    }
    bytes, err := json.MarshalIndent(output, "", " ")
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      fmt.Fprint(w,err.Error())
      return
    }
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Write(bytes)
  })
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	log.Fatal(http.ListenAndServe(":3050", r))
}
