package v1

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

const (
	ctxKeyUsers = iota
)

func (h *Handler) parseAuthHeader(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")

	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.tokenManager.Parse(headerParts[1])
}

func (h *Handler) usersIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := h.parseAuthHeader(r)
		if err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUsers, id)))
	})
}
