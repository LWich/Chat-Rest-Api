package handler

import (
	"encoding/json"
	"net/http"
)

func (h *HelloHandler) sayHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello")
	}
}
