package auth

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *AuthHandler) {
	mux.HandleFunc("POST /auth/register", h.Register)
}
