package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"go-key-value/internal/interfaces"
	"go-key-value/internal/redisRepository"
	"go-key-value/internal/memcachedRepository"

	"go-key-value/pkg/conf"
	"go-key-value/pkg/redis"
	"go-key-value/pkg/memcached"
	proto "go-key-value/pkg/keyvalue"

	"go-key-value/internal/http"
	localGRPC "go-key-value/internal/grpc"

	"github.com/gofiber/fiber/v3"
	"google.golang.org/grpc"
)

const db = "memcached"
const network = "grpc"

func main() {
	flagPath := flag.String("conf_path", "./pkg/conf/conf.json", "path to config")

	config, err := conf.ReadConf(flagPath)
	if err != nil {
		log.Fatal(err)
	}

	var repository interfaces.KeyValueRepositoryInterface

	switch db {
	case "redis":
		redisConn, err := redis.RedisConn(config)
		if err != nil {
			log.Fatal(err)
		}
	
		repository = redisrepository.NewRedisRepository(redisConn)
	case "memcached":
		memcachedConn, err := memcached.MemcachedConn(config)
		if err != nil {
			log.Fatal(err)
		}
	
		repository = memcachedRepository.NewMemcachedRepository(memcachedConn)
	default:
		log.Fatal(db)
	}

	switch network {
	case "http":
		router := fiber.New()
		http.NewAuthRouting(router, http.NewHandler(repository))
	
		log.Fatal(router.Listen(fmt.Sprintf(":%d", config.Main.HTTPPort)))
	case "grpc":
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Main.Host, config.Main.GRPCPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		proto.RegisterKeyvalueServer(s, localGRPC.NewAuthHandler(repository))
		log.Printf("server listening at %v", lis.Addr())
		
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	default:
		log.Fatal(network)
	}
}
