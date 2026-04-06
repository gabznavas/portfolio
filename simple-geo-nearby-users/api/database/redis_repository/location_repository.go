package redisrepository

import (
	"api/database/models"
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisLocationRepository interface {
	PutLocation(ctx context.Context, location models.Location) (success bool, err error)
}

type redisLocationRepositoryImpl struct {
	rdb *redis.Client
}

func NewRedisLocationRepository(rdb *redis.Client) RedisLocationRepository {
	return &redisLocationRepositoryImpl{rdb}
}

func (r *redisLocationRepositoryImpl) PutLocation(ctx context.Context, location models.Location) (success bool, err error) {
	success = true
	cmd := r.rdb.GeoAdd(ctx, "users", &redis.GeoLocation{
		Name:      location.Username,
		Longitude: location.Latitude,
		Latitude:  location.Longintude,
	})
	err = cmd.Err()
	success = err == nil
	return
}
