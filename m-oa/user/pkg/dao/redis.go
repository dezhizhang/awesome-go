package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var Rc *redis.Client

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       10,
	})
	//&RedisCache{
	//	rdb,
	//}

	_ := &RedisCache{
		rdb: rdb,
	}

}

func (rc *RedisCache) Put(ctx context.Context, key string, value string, expire time.Duration) error {
	return rc.rdb.Set(ctx, key, value, expire).Err()
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return rc.rdb.Get(ctx, key).Result()
}
