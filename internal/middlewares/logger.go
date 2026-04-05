package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now();
		next.ServeHTTP(w, req);
		log.Printf(
			"METHOD: %s | PATH: %s | DURATION: %v",
			req.Method,
			req.URL.Path,
			time.Since(start),
		);
	});
}
