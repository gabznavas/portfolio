package handlers

import (
	"api/database/models"
	redisrepository "api/database/redis_repository"
	handlers "api/handlers/dtos"
	helpers "api/handlers/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LocationHandler interface {
	CreateLocation(ctx *gin.Context)
	ListLocationByRange(ctx *gin.Context)
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

	username := "user:" + body.Username

	h.rou.PutOnlineUser(ctx, username)
	h.rlr.PutLocation(ctx, models.Location{
		Username:  username,
		Latitude:  body.Latitude,
		Longitude: body.Longitude,
	})
	ctx.JSON(http.StatusCreated, gin.H{
		"username": username,
		"lat":      body.Latitude,
		"long":     body.Longitude,
	})
}

func (h *locationHandlerImpl) ListLocationByRange(ctx *gin.Context) {
	latitude, err := helpers.ParseQueryFloat(ctx, "latitude")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	longitude, err := helpers.ParseQueryFloat(ctx, "longitude")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	radiusKm, err := helpers.ParseQueryFloat(ctx, "radiusKm")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// TODO: add services here

	locations, err := h.rlr.GetLocationsByPosition(ctx, latitude, longitude, &radiusKm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	removedPrefixRedix := []*models.Location{}
	for _, location := range locations {
		usernameWithoutUserPrefix := strings.Replace(location.Username, "user:", "", 1)
		location.Username = usernameWithoutUserPrefix
		removedPrefixRedix = append(removedPrefixRedix, location)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"locations": removedPrefixRedix,
	})
}
