package http

import (
	"net/http"

	v1 "github.com/LWich/chat-rest-api/internal/app/delivery/http/v1"
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
func New(service *service.Service, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		router:       mux.NewRouter(),
		service:      service,
		tokenManager: tokenManager,
	}
}

// ServeHTTP ...
func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.router.ServeHTTP(w, req)
}

// Init ...
func (h *Handler) Init() {
	h.initApi()
}

func (h *Handler) initApi() {
	v1 := v1.New(h.tokenManager, h.service)
	api := h.router.PathPrefix("/api").Subrouter()
	{
		v1.Init(api)
	}
}
