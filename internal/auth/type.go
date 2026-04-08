package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AuthHandler struct {
	db          *sqlx.DB
	redisClient *redis.Client
}

func NewHandler(db *sqlx.DB, redisClient *redis.Client) *AuthHandler {
	return &AuthHandler{db, redisClient}
}

type RegisterDto struct {
	Name      string `json:"name"     validate:"required,min=3,max=255"`
	Age       int    `json:"age"      validate:"required,gte=0,lte=130"`
	Email     string `json:"email"    validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=73,password"`
}

type LoginDto struct {
	Email     string `json:"email"    validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=73,password"`
}
