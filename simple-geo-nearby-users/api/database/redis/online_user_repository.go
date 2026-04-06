package redisrepository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisOnlineUsers interface {
	PutOnlineUser(ctx context.Context, username string) (success bool, err error)
}

type redisOnlineUsersImpl struct {
	rdb *redis.Client
}

func NewRedisOnlineUsers(rdb *redis.Client) RedisOnlineUsers {
	return &redisLocationRepositoryImpl{rdb}
}

func (r *redisLocationRepositoryImpl) PutOnlineUser(ctx context.Context, username string) (success bool, err error) {
	cmd := r.rdb.SAdd(ctx, "online_users", username)
	err = cmd.Err()
	success = err == nil
	return
}
