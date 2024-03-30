package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	redisrepository "redis/internal/redisRepository"
	"redis/pkg/conf"
	"redis/pkg/redis"

	auth "redis/internal/auth"

	"github.com/gofiber/fiber/v3"
)

func main() {
	flagPath := flag.String("conf_path", "./pkg/conf/conf.json", "path to config")

	config, err := conf.ReadConf(flagPath)
	if err != nil {
		log.Fatal(err)
	}

	redisConn := redis.RedisConn(config)
	
	if err := redisConn.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	redisDB := redisrepository.NewRedisRepository(redisConn)
	AuthHandler := auth.NewAuthHandler(redisDB)


	router := fiber.New()
	auth.NewAuthRouting(router, AuthHandler)

	log.Fatal(router.Listen(fmt.Sprintf(":%d", config.Main.Port)))
}
