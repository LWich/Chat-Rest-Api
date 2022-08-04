package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LWich/chat-rest-api/internal/app/model"
	"github.com/gorilla/mux"
)

func (h *Handler) initUsersRoutes(router *mux.Router) {
	users := router.PathPrefix("/users").Subrouter()
	{
		users.HandleFunc("/signup", h.userSignUp())
	}
}

func (h *Handler) userSignUp() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		fmt.Println(r.Body)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := h.store.User().Create(u); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}

		u.Sanitize()

		h.respond(w, r, http.StatusOK, u)
	}
}
