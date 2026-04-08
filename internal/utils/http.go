package utils

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "apliction-json");
	w.WriteHeader(status);
	json.NewEncoder(w).Encode(data);
}

func WriteCookie(w http.ResponseWriter, n string, v string) {
	http.SetCookie(w, &http.Cookie{
		Name: n,
		Value: v,
		Path: "/",
		MaxAge: 0,
		HttpOnly:true,
	})
}
