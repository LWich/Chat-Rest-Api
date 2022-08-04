package v1

import (
	"encoding/json"
	"net/http"

	"github.com/LWich/chat-rest-api/internal/app/store"
	"github.com/gorilla/mux"
)

// Handler ...
type Handler struct {
	store  *store.Store
	router *mux.Router
}

// New ...
func New(store *store.Store) *Handler {
	return &Handler{
		store:  store,
		router: mux.NewRouter(),
	}
}

// Init ...
func (h *Handler) Init() {
	h.router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		h.respond(w, r, 200, map[string]string{"msg": "Hello"})
	}).Methods(http.MethodGet)

	v1 := h.router.PathPrefix("/v1").Subrouter()
	{
		v1.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			h.respond(w, r, 200, map[string]string{"msg": "Hello"})
		}).Methods(http.MethodGet)
		h.initUsersRoutes(v1)
	}
}

// ServeHTTP ...
func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.router.ServeHTTP(w, req)
}

func (h *Handler) error(w http.ResponseWriter, req *http.Request, code int, err error) {
	h.respond(w, req, code, map[string]string{"Error": err.Error()})
}

func (h *Handler) respond(w http.ResponseWriter, req *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
