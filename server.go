package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pace-noge/golang-gin-poc/controllers"
	"github.com/pace-noge/golang-gin-poc/services"
)

var (
	videoService    services.VideoService       = services.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(201, videoController.Save(ctx))
	})

	server.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "healthy.",
		})
	})

	server.Run(":8080")
}
