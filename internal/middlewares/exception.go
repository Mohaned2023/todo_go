package middlewares

import (
	"encoding/json"
	"fmt"

	"net/http"
	"todo/internal/types"
	"todo/pkg/logger"
)


func handleError(w http.ResponseWriter, err any) {
	var appError types.AppError;
	if e, ok := err.(types.AppError); ok {
		appError = e;
	} else {
		logger.Err(fmt.Errorf("%v", err));
		appError = types.AppError{
			Status: http.StatusInternalServerError,
			Message: "An unexpected error occurred",
		}
	}
	w.Header().Set("Content-Type", "apliction-json");
	w.WriteHeader(appError.Status);
	json.NewEncoder(w).Encode(appError);
}

func EceptionHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				handleError(w, r);
			}
		}();
		next.ServeHTTP(w, req);
	});
}
