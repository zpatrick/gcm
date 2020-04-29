package main

import (
	"log"

	"github.com/zpatrick/gcm/example/cfg"
	"github.com/zpatrick/gcm/example/redis"
)

func main() {
	cm, err := cfg.Load()
	if err != nil {
		log.Fatal(err)
	}

	redisClient := redis.NewClient(redis.ClientConfig{
		Host: cm.MustString(cfg.KeyRedisHost),
		Port: cm.MustInt(cfg.KeyRedisPort),
	})

	if err := redisClient.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")
}
