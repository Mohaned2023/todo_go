package server

import (
	"fmt"
	"log/slog"
	"os"
	"todo/internal/storage"
)

func getEnv(key string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	panic(fmt.Sprintf("You must add %s to your ENV!", key))
}

func configSlog(env string) {
	opt := slog.HandlerOptions{}
	if env == "prod" {
		opt.Level = slog.LevelInfo
	} else {
		opt.Level = slog.LevelDebug
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &opt))
	slog.SetDefault(logger)
}

func lodConfig() Config {
	configSlog(getEnv("ENV"))
	
	return Config{
		AppPort: getEnv("APP_PORT"),
		DB: storage.InitDB(getEnv("DATABASE_URL")),
		RedisClient: storage.InitRedis(getEnv("REDIS_HOST")),
		CORSOrigin: getEnv("CORS_ORIGIN"),
	}
}
