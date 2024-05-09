package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var Ctx = context.Background()

func CreateClient(db int) *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr: os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASSWORD"),
		DB: db,
	})

	return redis
}