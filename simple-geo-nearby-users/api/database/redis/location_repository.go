package redisrepository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisLocationRepository interface {
	PutLocation(ctx context.Context, username string, latitude, longitude float64) (success bool, err error)
}

type redisLocationRepositoryImpl struct {
	rdb *redis.Client
}

func NewRedisLocationRepository(rdb *redis.Client) RedisLocationRepository {
	return &redisLocationRepositoryImpl{rdb}
}

func (r *redisLocationRepositoryImpl) PutLocation(ctx context.Context, username string, latitude, longitude float64) (success bool, err error) {
	success = true
	cmd := r.rdb.HSet(ctx, username, map[string]interface{}{
		"lat":  latitude,
		"long": longitude,
	})
	err = cmd.Err()
	success = err == nil
	return
}
