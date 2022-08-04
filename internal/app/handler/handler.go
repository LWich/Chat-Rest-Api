package handler

import (
	"net/http"

	"github.com/LWich/chat-rest-api/internal/app/store"
	"github.com/gorilla/mux"
)

// HelloHandler ...
type HelloHandler struct {
	router *mux.Router
	store  *store.Store
}

// NewHelloHandler ...
func NewHelloHandler(store *store.Store) *HelloHandler {
	h := &HelloHandler{
		router: mux.NewRouter(),
		store:  store,
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
