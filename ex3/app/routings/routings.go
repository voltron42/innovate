package routings

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Hello struct {
	path string
}

func (h *Hello) Path() string {
	return h.path
}

func (h *Hello) Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h2>Hello World.</h2>"))
}

type HelloQuery string

func (h *HelloQuery) Path() string {
	return string(*h)
}

func (h *HelloQuery) Handler(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query["name"]
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if len(name) > 0 {
		fmt.Fprintf(w, "<h2>Hello %v.</h2>", strings.Join(name, " "))
	} else {
		w.Write([]byte("<h2>Hello World.</h2>"))
	}
}

type helloPath string

func (h *helloPath) Path() string {
	return string(*h)
}

func (h *helloPath) Handler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h2>Hello %v.</h2>", vars["name"])
}

func NewHelloPath(root string) *helloPath {
	path := helloPath(fmt.Sprintf("/%v/{name}", root))
	return &path
}
