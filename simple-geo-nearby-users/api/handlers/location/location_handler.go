package handlers

import (
	redisrepository "api/database/redis"
	handlers "api/handlers/dtos"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LocationHandler interface {
	CreateLocation(ctx *gin.Context)
}

type locationHandlerImpl struct {
	rlr redisrepository.RedisLocationRepository
	rou redisrepository.RedisOnlineUsers
}

func NewLocationHandler(
	rlr redisrepository.RedisLocationRepository,
	rou redisrepository.RedisOnlineUsers,
) LocationHandler {
	return &locationHandlerImpl{rlr, rou}
}

func (h *locationHandlerImpl) CreateLocation(ctx *gin.Context) {
	var body handlers.CreateLocationRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	username := fmt.Sprintf("users:%s", body.Username)

	h.rou.PutOnlineUser(ctx, username)
	h.rlr.PutLocation(ctx, username, body.Latitude, body.Longitude)

	ctx.JSON(http.StatusCreated, gin.H{
		"username": username,
		"lat":      body.Latitude,
		"long":     body.Longitude,
	})
}
