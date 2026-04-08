package session

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
	"todo/internal/user"

	"github.com/redis/go-redis/v9"
)

const REDIS_PERF_KEY string = "session"
const REDIS_TTL time.Duration = time.Hour

func MakeSession(ctx context.Context, redisClient *redis.Client, u *user.User) string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	salt := base64.RawStdEncoding.EncodeToString(b)
	key := fmt.Sprintf("%s.%s.%s", u.Email, u.Password, salt)
	h := sha256.New()
	h.Write([]byte(key))
	sid := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	data, _ := json.Marshal(u)

    redisClient.Set(
		ctx,
		fmt.Sprintf("%s:%s", REDIS_PERF_KEY, sid),
		data,
		REDIS_TTL,
	)

	return sid
}


func GetAndUpdateTTL(ctx context.Context, redisClient *redis.Client, sid string) (*user.User, error) {
	val, err := redisClient.GetEx(
		ctx,
		fmt.Sprintf("%s:%s", REDIS_PERF_KEY, sid),
		REDIS_TTL,
	).Result()
	if err != nil {
		return nil, err
	}

	var u user.User;
	json.Unmarshal([]byte(val), &u)
	return &u, nil
}

