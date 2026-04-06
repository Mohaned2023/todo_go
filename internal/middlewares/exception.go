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
	if e, ok := err.(types.AppErrorTypes); ok {
		appError.Map(e);
	} else {
		logger.Err(fmt.Errorf("%v", err));
		appError.Map(types.InteralServerError)
	}
	w.Header().Set("Content-Type", "apliction-json");
	w.WriteHeader(appError.Status);
	json.NewEncoder(w).Encode(appError);
}

func ExceptionHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				handleError(w, r);
			}
		}();
		next.ServeHTTP(w, req);
	});
}
