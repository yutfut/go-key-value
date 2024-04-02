package main

import (
	"flag"
	"fmt"
	"log"

	"go-key-value/internal/redisRepository"
	"go-key-value/internal/memcachedRepository"

	"go-key-value/pkg/conf"
	"go-key-value/pkg/redis"
	"go-key-value/pkg/memcached"
	"go-key-value/pkg/riak"

	"go-key-value/internal/auth"

	"github.com/gofiber/fiber/v3"
)

const db = "memcached"

func main() {
	flagPath := flag.String("conf_path", "./pkg/conf/conf.json", "path to config")

	config, err := conf.ReadConf(flagPath)
	if err != nil {
		log.Fatal(err)
	}

	var AuthHandler auth.AuthHandlerInterface

	switch db{
	case "redis":
		redisConn, err := redis.RedisConn(config)
		if err != nil {
			log.Fatal(err)
		}
	
		redisDB := redisrepository.NewRedisRepository(redisConn)
		AuthHandler = auth.NewAuthHandler(redisDB)
	case "memcached":
		memcachedConn, err := memcached.MemcachedConn(config)
		if err != nil {
			log.Fatal(err)
		}
	
		memcachedDB := memcachedRepository.NewMemcachedRepository(memcachedConn)
		AuthHandler = auth.NewAuthHandler(memcachedDB)
	default:
		log.Fatal(db)
	}

	riak.RiakConn()

	router := fiber.New()
	auth.NewAuthRouting(router, AuthHandler)

	log.Fatal(router.Listen(fmt.Sprintf(":%d", config.Main.Port)))
}
