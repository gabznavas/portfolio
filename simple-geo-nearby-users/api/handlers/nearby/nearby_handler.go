package handlers

import (
	"fmt"
	"net/http"

	dtos "api/handlers/dtos"
	helpers "api/handlers/helpers"

	"github.com/gin-gonic/gin"
)

type NearbyHandler interface {
	ListNearbyByPosition(ctx *gin.Context)
}

type nearbyHandlerImpl struct{}

func NewNearbyHandler() NearbyHandler {
	return &nearbyHandlerImpl{}
}

func (h *nearbyHandlerImpl) ListNearbyByPosition(ctx *gin.Context) {
	latitude, err := helpers.ParseQueryFloat(ctx, "latitude")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	longitude, err := helpers.ParseQueryFloat(ctx, "longitude")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%f %f", latitude, longitude)
	var fakePositions = []dtos.ListNearbyByPositionResponse{
		{
			Username:  "mario1",
			Latitude:  -22.222,
			Longitude: -51.222,
		}, {
			Username:  "mario2",
			Latitude:  -22.222,
			Longitude: -51.222,
		},
	}
	ctx.JSON(http.StatusOK, fakePositions)
}
