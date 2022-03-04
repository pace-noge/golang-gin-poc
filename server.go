package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pace-noge/golang-gin-poc/controllers"
	"github.com/pace-noge/golang-gin-poc/middlewares"
	"github.com/pace-noge/golang-gin-poc/services"
	ginDump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    services.VideoService       = services.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

// func setupLogOutput() {
// 	// create gin.log file
// 	f, _ := os.Create("gin.log")

// 	// write log to file and stdout
// 	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
// }

func main() {
	// setupLogOutput()

	server := gin.New()

	server.Use(
		gin.Recovery(),
		middlewares.BasicAuth(),
		middlewares.Logger(),
		ginDump.Dump(),
	)

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
