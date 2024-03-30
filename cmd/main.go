package main

import (
	"flag"
	"fmt"
	"log"

	"redis/internal/redisRepository"
	// "redis/internal/memcachedRepository"

	"redis/pkg/conf"
	"redis/pkg/redis"
	// "redis/pkg/memcached"

	auth "redis/internal/auth"

	"github.com/gofiber/fiber/v3"
)

func main() {
	flagPath := flag.String("conf_path", "./pkg/conf/conf.json", "path to config")

	config, err := conf.ReadConf(flagPath)
	if err != nil {
		log.Fatal(err)
	}

	redisConn, err := redis.RedisConn(config)
	if err != nil {
		log.Fatal(err)
	}

	redisDB := redisrepository.NewRedisRepository(redisConn)
	AuthHandler := auth.NewAuthHandler(redisDB)

	// memcachedConn, err := memcached.MemcachedConn(config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// memcachedDB := memcachedRepository.NewMemcachedRepository(memcachedConn)
	// AuthHandler := auth.NewAuthHandler(memcachedDB)


	router := fiber.New()
	auth.NewAuthRouting(router, AuthHandler)

	log.Fatal(router.Listen(fmt.Sprintf(":%d", config.Main.Port)))
}
