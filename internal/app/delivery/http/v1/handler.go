package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LWich/chat-rest-api/internal/app/config"
	"github.com/LWich/chat-rest-api/internal/app/store"
	"github.com/LWich/chat-rest-api/internal/app/tokenmanager"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// Handler ...
type Handler struct {
	store           *store.Store
	router          *mux.Router
	sessionStore    sessions.Store
	tokenManager    *tokenmanager.Manager
	tokenmanagerCfg config.AuthConfig
}

// New ...
func New(store *store.Store,
	sessionStore sessions.Store,
	managerCfg config.AuthConfig,
) *Handler {
	return &Handler{
		store:           store,
		router:          mux.NewRouter(),
		tokenmanagerCfg: managerCfg,
		tokenManager:    tokenmanager.NewManager(managerCfg.SigninKey),
		sessionStore:    sessionStore,
	}
}

// Init ...
func (h *Handler) Init() {
	v1 := h.router.PathPrefix("/v1").Subrouter()
	{
		v1.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(w.Header().Get("X-Auth-Token"))
		})
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
