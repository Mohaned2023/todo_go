package server

import (
	"net/http"
	"todo/internal/auth"
	"todo/internal/middlewares"
)

func server(cfg Config)  http.Handler {
	rootMux := http.NewServeMux()
	apiV1Mux := http.NewServeMux()

	auth.RegisterRoutes(apiV1Mux, auth.NewHandler(cfg.DB, cfg.RedisClient))

	http.Handle("/v1/", http.StripPrefix("/v1", apiV1Mux))

	muxWapper := middlewares.ExceptionHandler(rootMux)
	muxWapper = middlewares.EnableCORS(muxWapper, cfg.CORSOrigin)
	muxWapper = middlewares.Logger(muxWapper)

	return muxWapper
}
