package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	redisrepository "api/database/redis_repository"
	locationHandler "api/handlers/location"
	nearbyHandler "api/handlers/nearby"
)

const (
	HTTP_SERVER_PORT = "HTTP_SERVER_PORT"
	REDIS_ADDR       = "REDIS_ADDR"
)

func loadEnvs() map[string]string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error on load .env")
	}

	var (
		httpServerPort = os.Getenv(HTTP_SERVER_PORT)
		redisAddress   = os.Getenv(REDIS_ADDR)
	)
	if httpServerPort == "" {
		panic("needed " + HTTP_SERVER_PORT + " on .env")
	}
	if redisAddress == "" {
		panic("needed " + REDIS_ADDR + " on .env")
	}

	return map[string]string{
		HTTP_SERVER_PORT: httpServerPort,
		REDIS_ADDR:       redisAddress,
	}

}

func main() {
	var envs map[string]string = loadEnvs()

	router := gin.Default()
	router.Use(cors.Default())

	rdb := redis.NewClient(&redis.Options{
		Addr: envs["REDIS_ADDR"],
	})

	locationRepo := redisrepository.NewRedisLocationRepository(rdb)
	onlineUsersRepo := redisrepository.NewRedisOnlineUsers(rdb)

	locationHdl := locationHandler.NewLocationHandler(
		locationRepo,
		onlineUsersRepo,
	)
	nearbyHdl := nearbyHandler.NewNearbyHandler()

	router.POST("/location", locationHdl.CreateLocation)
	router.GET("/nearby", nearbyHdl.ListNearbyByPosition)

	router.Run(fmt.Sprintf(":%s", envs[HTTP_SERVER_PORT]))
}
