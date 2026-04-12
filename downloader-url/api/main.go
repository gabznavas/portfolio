package main

import (
	"api/services"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Urls struct {
	Urls []string `json:"urls"`
}

func initStorage() {
	err := os.MkdirAll("storage", os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
}

func DownloadByUrlHandlerfunc(ctx *gin.Context) {
	var body Urls
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	id := uuid.New()
	urls := []*url.URL{}

	for _, rawUrl := range body.Urls {
		url, err := url.ParseRequestURI(rawUrl)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		urls = append(urls, url)
	}

	services.ExecDownload(urls, id)

	ctx.JSON(http.StatusCreated, gin.H{
		"requestId": string(id.String()),
	})
}

func main() {
	initStorage()

	route := gin.Default()
	route.POST("/", DownloadByUrlHandlerfunc)
	route.Run(":8081")
}
