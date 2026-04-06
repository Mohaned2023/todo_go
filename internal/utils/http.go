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
