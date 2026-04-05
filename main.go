package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"todo/internal/handlers"
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
	storage.InitDB(dbUrl);
	defer storage.DBConn.Close();

	mux := http.NewServeMux();
	mux.HandleFunc("POST /v1/auth/register", handlers.AuthRegister);

	wapperMux := middlewares.Logger(mux);
	wapperMux = middlewares.EceptionHandler(wapperMux);
	
	port := os.Getenv("APP_PORT");
	if port == "" {
		panic("You must set APP_PORT!");
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), wapperMux);
}
