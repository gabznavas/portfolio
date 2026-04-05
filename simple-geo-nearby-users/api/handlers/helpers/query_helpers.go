package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseQueryFloat(ctx *gin.Context, key string) (float64, error) {
	value := ctx.Query(key)
	if value == "" {
		return 0, fmt.Errorf("%s is required", key)
	}

	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s", key)
	}

	return f, nil
}
