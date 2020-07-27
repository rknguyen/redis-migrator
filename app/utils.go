package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/huyng12/redis-migrate-db/config"
)

func MakeConnection(ctx context.Context, conn *config.RedisConnection) *redis.Client {
	addr := fmt.Sprintf("%s:%d", conn.Host, conn.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conn.Password,
		DB:       conn.DatabaseIndex,
	})

	// test connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to %s ", addr), err)
	}

	log.Printf("Connected to Redis: addr=%s, db=%d", addr, conn.DatabaseIndex)
	return rdb
}

func TearDownConnection(client *redis.Client) {
	err := client.Close()
	if err != nil {
		log.Print("Failed to close the connection", err)
	} else {
		log.Print("Closed the connection")
	}
}
