package routings

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"../model"
	"io/ioutil"
	"encoding/json"
)

type Hello struct {
	path string
}

func NewHello(path string) *Hello {
	return &Hello{path: path}
}

func (h *Hello) Path() string {
	return h.path
}

func (h *Hello) Method() model.Method {
	return model.GET
}

func (h *Hello) Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h2>Hello World.</h2>"))
}

type HelloQuery string

func (h *HelloQuery) Path() string {
	return string(*h)
}

func (h *HelloQuery) Method() model.Method {
	return model.GET
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

func (h *helloPath) Method() model.Method {
	return model.GET
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

type HelloJson string

func (h *HelloJson) Path() string {
	return string(*h)
}

func (h *HelloJson) Method() model.Method {
	return model.POST
}

func (h *HelloJson) Handler(w http.ResponseWriter, req *http.Request) {
	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	name := Name{}
	err = json.Unmarshal(bytes, &name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h2>Hello %v.</h2>", name.String())
}

type Name struct {
	First string `json:"first"`
	Middle string `json:"middle,omitempty"`
	Last string `json:"last"`
}

func (n Name) String() string {
	list := []string{
		 n.First,
	}
	if len(n.Middle) > 0 {
		list = append(list, n.Middle)
	}
	list = append(list, n.Last)
	return strings.Join(list, " ")
}
