package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)


func Run() {
	godotenv.Load();
	config := lodConfig()
	serverMux := server(config)
	ser := http.Server{
		Addr: ":" + config.AppPort,
		Handler: serverMux,
	}

	defer config.DB.Close()
	defer config.RedisClient.Close()
	go func() {
		slog.Info("Running server.", "port", config.AppPort)
		if err := ser.ListenAndServe(); err != nil {
			slog.Error("Server error", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown the server...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ser.Shutdown(ctx)
}
