package server

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	AppPort     string
	DB          *sqlx.DB
	RedisClient *redis.Client
	CORSOrigin  string
}
