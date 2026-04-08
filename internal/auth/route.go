package auth

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *AuthHandler) {
	mux.HandleFunc("POST /auth/register", h.Register)
	mux.HandleFunc("POST /auth/login", h.Login)
}
