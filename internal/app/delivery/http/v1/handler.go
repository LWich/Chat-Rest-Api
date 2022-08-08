package v1

import (
	"encoding/json"
	"net/http"

	"github.com/LWich/chat-rest-api/internal/app/service"
	"github.com/LWich/chat-rest-api/pkg/auth"
	"github.com/gorilla/mux"
)

// Handler ...
type Handler struct {
	router       *mux.Router
	tokenManager auth.TokenManager
	service      *service.Service
}

// New ...
func New(tokenManager auth.TokenManager, service *service.Service) *Handler {
	return &Handler{
		tokenManager: tokenManager,
		service:      service,
	}
}

// Init ...
func (h *Handler) Init(api *mux.Router) {
	h.router = api
	v1 := api.PathPrefix("/v1").Subrouter()
	{
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
