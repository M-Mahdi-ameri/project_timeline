package config

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx = context.Background()

func Initredis() {
	addr := os.Getenv("REDIS_ADDR")
	pass := os.Getenv("REDIS_PASS")

	if addr == "" {
		addr = "127.0.0.1:6379"
	}

	RDB = redis.NewClient(&redis.Options{
		Addr: addr, DB: 0, Password: pass,
	})
	if _, err := RDB.Ping(ctx).Result(); err != nil {
		log.Fatalf("falid to connect to redis: %v", err)

	}
	log.Println("connect to redis succesfully")
}
