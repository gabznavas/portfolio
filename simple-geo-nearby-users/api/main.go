package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	locationHandler "api/handlers/location"
	nearbyHandler "api/handlers/nearby"
)

const (
	HTTP_SERVER_PORT = "HTTP_SERVER_PORT"
)

func loadEnvs() map[string]string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error on load .env")
	}

	var (
		httpServerPort = os.Getenv(HTTP_SERVER_PORT)
	)
	if httpServerPort == "" {
		panic("needed " + HTTP_SERVER_PORT + " on .env")
	}

	return map[string]string{
		HTTP_SERVER_PORT: httpServerPort,
	}

}

func main() {
	var envs map[string]string = loadEnvs()

	router := gin.Default()
	router.Use(cors.Default())

	locationHdl := locationHandler.NewLocationHandler()
	nearbyHdl := nearbyHandler.NewNearbyHandler()

	router.POST("/location", locationHdl.CreateLocation)
	router.GET("/nearby", nearbyHdl.ListNearbyByPosition)

	router.Run(fmt.Sprintf(":%s", envs[HTTP_SERVER_PORT]))
}
