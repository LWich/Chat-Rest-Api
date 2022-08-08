package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LWich/chat-rest-api/internal/app/service"
	"github.com/gorilla/mux"
)

// Tokens ...
type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func (h *Handler) initUsersRoutes(router *mux.Router) {
	users := router.PathPrefix("/users").Subrouter()
	{
		users.HandleFunc("/signup", h.userSignUp())
		users.HandleFunc("/signin", h.userSignIn())
		users.HandleFunc("/auth/refresh", h.userRefresh())
	}
}

func (h *Handler) userRefresh() http.HandlerFunc {
	type request struct {
		Token string `json:"token"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}

		tokens, err := h.service.Users.RefreshTokens(req.Token)
		if err != nil {
			var code int
			if err == service.ErrFailedToCreateTokens {
				code = http.StatusInternalServerError
			}

			code = http.StatusBadRequest

			h.error(w, r, code, err)
			return
		}

		h.respond(w, r, http.StatusOK, tokens)
	}
}

func (h *Handler) userSignIn() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}

		tokens, err := h.service.Users.SignIn(service.UsersSignInInput{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			var code int
			if err == service.ErrFailedToCreateTokens {
				code = http.StatusInternalServerError
			}

			code = http.StatusUnauthorized

			h.error(w, r, code, err)
			return
		}

		h.respond(w, r, http.StatusOK, tokens)
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

		u, err := h.service.Users.SignUp(service.UsersSignUpInput{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		h.respond(w, r, http.StatusOK, u)
	}
}
