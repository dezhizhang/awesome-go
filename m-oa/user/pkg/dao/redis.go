package dao

import (
	"awesome-go/m-oa/user/config"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

type RedisCache struct {
	rdb *redis.Client
}

var Rc *RedisCache

func init() {
	rdb := redis.NewClient(config.C.ReadRedisConfig())
	zap.L().Info("连接redis成功")
	Rc = &RedisCache{
		rdb: rdb,
	}

}

func (rc *RedisCache) Put(ctx context.Context, key string, value string, expire time.Duration) error {
	return rc.rdb.Set(ctx, key, value, expire).Err()
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return rc.rdb.Get(ctx, key).Result()
}
