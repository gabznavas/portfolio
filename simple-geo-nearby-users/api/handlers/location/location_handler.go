package handlers

import (
	handlers "api/handlers/dtos"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	fmt.Printf("%s %f %f", body.Username, body.Latitude, body.Longitude)

	ctx.JSON(http.StatusNoContent, nil)
}
