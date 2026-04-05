package handlers

import (
	handlers "api/handlers/dtos"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type LocationHandler interface {
	CreateLocation(ctx *gin.Context)
}

type locationHandlerImpl struct{}

func NewLocationHandler() LocationHandler {
	return &locationHandlerImpl{}
}

func (h *locationHandlerImpl) CreateLocation(ctx *gin.Context) {
	var body handlers.CreateLocationRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	username := fmt.Sprintf("users:%s", body.Username)

	cmd := rdb.SAdd(ctx, "online_users", username)
	err := cmd.Err()
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	cmd = rdb.HSet(ctx, username, map[string]interface{}{
		"lat":  body.Latitude,
		"long": body.Longitude,
	})
	err = cmd.Err()
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"username": username,
		"lat":      body.Latitude,
		"long":     body.Longitude,
	})
}
