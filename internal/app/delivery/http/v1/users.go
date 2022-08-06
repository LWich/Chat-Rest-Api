package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/LWich/chat-rest-api/internal/app/model"
	"github.com/gorilla/mux"
)

var (
	sessionName = "refresh-token"
)

var (
	ErrPasswordOrEmailIncorrect = errors.New("password or email incorrect")
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

		u, err := h.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparerPassword(req.Password) {
			h.error(w, r, http.StatusUnauthorized, ErrPasswordOrEmailIncorrect)
			return
		}

		tokens, err := h.createSession(w, r, u.Id)
		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		h.respond(w, r, http.StatusOK, tokens)
	}
}

func (h *Handler) createSession(w http.ResponseWriter, r *http.Request, userId int) (*Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = h.tokenManager.NewJWT(userId, time.Duration(h.tokenmanagerCfg.AccessTokenTTL))
	if err != nil {
		return nil, err
	}

	res.RefreshToken = h.tokenManager.NewRefreshToken()

	session, err := h.sessionStore.Get(r, sessionName)
	if err != nil {
		return nil, err
	}

	session.Values["refreshToken"] = res.RefreshToken
	session.Options.MaxAge = h.tokenmanagerCfg.RefreshTokenTTL

	if err := session.Save(r, w); err != nil {
		return nil, err
	}

	return &res, h.store.User().SetRefreshTokenBySession(userId, session)
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
