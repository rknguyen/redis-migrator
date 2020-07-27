package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func GetAllKeys(ctx context.Context, client *redis.Client) ([]string, error) {
	var cursor uint64
	var allKeys []string

	for {
		var err error
		var keys []string
		if keys, cursor, err = client.Scan(ctx, cursor, "", 50).Result(); err != nil {
			return nil, err
		}

		allKeys = append(allKeys, keys...)

		if cursor == 0 {
			break
		}
	}

	return allKeys, nil
}
