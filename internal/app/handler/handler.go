package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HelloHandler ...
type HelloHandler struct {
	router *mux.Router
}

// NewHelloHandler ...
func NewHelloHandler() *HelloHandler {
	h := &HelloHandler{
		router: mux.NewRouter(),
	}

	h.configureRouter()

	return h
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.router.ServeHTTP(w, req)
}

func (h *HelloHandler) configureRouter() {
	h.router.HandleFunc("/sayhello", h.sayHello()).Methods(http.MethodGet)
}
