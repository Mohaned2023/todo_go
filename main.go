package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"todo/internal/auth"
	"todo/internal/middlewares"
	"todo/internal/storage"
)

func main() {
	err := godotenv.Load();
	if err != nil {
		panic("Can not load the env file!");
	}

	dbUrl := os.Getenv("DATABASE_URL");
	if dbUrl == "" {
		panic("You must set DATABASE_URL!");
	}
	db := storage.InitDB(dbUrl);
	defer db.Close();

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		panic("You must set REDIS_HOST!")
	}
	redisClient := storage.InitRedis(redisHost)
	defer redisClient.Close();

	rootMux := http.NewServeMux();
	apiV1Mux := http.NewServeMux();

	authHandler := auth.NewHandler(db, redisClient)
	auth.RegisterRoutes(apiV1Mux, authHandler)
	
	rootMux.Handle("/v1/", http.StripPrefix("/v1", apiV1Mux))

	wapperMux := middlewares.Logger(rootMux);
	wapperMux = middlewares.ExceptionHandler(wapperMux);
	
	port := os.Getenv("APP_PORT");
	if port == "" {
		panic("You must set APP_PORT!");
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), wapperMux);
}
