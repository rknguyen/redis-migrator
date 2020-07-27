package redis

import (
	"context"
	"log"
	"time"

	"github.com/gammazero/workerpool"

	"github.com/go-redis/redis/v8"
	"github.com/huyng12/redis-migrate-db/config"
)

type RedisMigrator struct {
	cfg config.Config
}

func NewRedisMigrator(cfg config.Config) *RedisMigrator {
	return &RedisMigrator{cfg: cfg}
}

func (rm *RedisMigrator) Migrate(ctx context.Context, src *redis.Client, dst *redis.Client) {
	keys, err := GetAllKeys(ctx, src)
	if err != nil {
		log.Fatal(err)
	}

	wp := workerpool.New(rm.cfg.ThreadNum)
	log.Printf("Created worker pool, thread num: %d", rm.cfg.ThreadNum)

	keySize := len(keys)
	log.Printf("Found %d keys", keySize)

	finished := 0
	for _, key := range keys {
		key := key
		wp.Submit(func() {
			var val string
			var ttl time.Duration
			if ttl, err = src.PTTL(ctx, key).Result(); err != nil {
				log.Fatal(err)
			}
			if val, err = src.Dump(ctx, key).Result(); err != nil {
				log.Fatal(err)
			}

			if err = dst.Del(ctx, key).Err(); err != nil {
				log.Fatal(err)
			}
			if err = dst.Restore(ctx, key, ttl, val).Err(); err != nil {
				log.Fatal(err)
			}

			finished++
			log.Printf("Setted %d/%d keys", finished, keySize)
		})
	}

	wp.StopWait()
}
