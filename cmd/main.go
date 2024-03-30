package main

import (
	"context"
	"log"

	redisrepository "redis/internal/redisRepository"
	"redis/pkg/redis"

	auth "redis/internal/auth"

	"github.com/gofiber/fiber/v3"
)

func main() {
	router := fiber.New()

	redisConn := redis.RedisConn()
	
	if err := redisConn.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	redisDB := redisrepository.NewRedisRepository(redisConn)
	AuthHandler := auth.NewAuthHandler(redisDB)

	auth.NewAuthRouting(router, AuthHandler)

	log.Fatal(router.Listen(":8000"))
}
