package main

import (
	"log"

	"github.com/zpatrick/gcm/example/config"
	"github.com/zpatrick/gcm/example/redis"
)

func main() {
	cm, err := config.Mux()
	if err != nil {
		log.Fatal(err)
	}

	redisClient := redis.NewClient(redis.ClientConfig{
		Host: cm.MustString(config.KeyRedisHost),
		Port: cm.MustInt(config.KeyRedisPort),
	})

	if err := redisClient.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")
}
