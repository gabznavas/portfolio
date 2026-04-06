package redisrepository

import (
	"api/database/models"
	"context"

	"github.com/redis/go-redis/v9"
)

const Key = "online_users"

type RedisOnlineUsers interface {
	PutOnlineUser(ctx context.Context, username models.OnlineUser) (success bool, err error)
	ListOnlineUsers(ctx context.Context) (onlineUsers []models.OnlineUser, err error)
}

type redisOnlineUsersImpl struct {
	rdb *redis.Client
}

func NewRedisOnlineUsers(rdb *redis.Client) RedisOnlineUsers {
	return &redisLocationRepositoryImpl{rdb}
}

func (r *redisLocationRepositoryImpl) PutOnlineUser(ctx context.Context, onlineUser models.OnlineUser) (success bool, err error) {
	cmd := r.rdb.SAdd(ctx, Key, onlineUser)
	err = cmd.Err()
	success = err == nil
	return
}

func (r *redisLocationRepositoryImpl) ListOnlineUsers(ctx context.Context) (onlineUsers []models.OnlineUser, err error) {
	cmd := r.rdb.SMembers(ctx, "online_users")
	onlineUsers, err = cmd.Result()
	return
}
