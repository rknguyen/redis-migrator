package main

import (
	"context"
	"log"

	RM "github.com/huyng12/redis-migrate-db/redis"

	"github.com/go-redis/redis/v8"

	"github.com/huyng12/redis-migrate-db/config"
)

func main() {
	cfg := config.LoadConfig()
	ctx := context.Background()

	// make sure that both two connections are valid
	src := MakeConnection(ctx, &cfg.Src)
	dst := MakeConnection(ctx, &cfg.Dst)

	// create new redis migrator
	rm := RM.NewRedisMigrator(*cfg)

	rm.Migrate(ctx, src, dst)
	log.Println("Migration is finished!")

	cleanup(src, dst)
}

func cleanup(src *redis.Client, dst *redis.Client) {
	TearDownConnection(src)
	TearDownConnection(dst)
}
