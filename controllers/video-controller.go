package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pace-noge/golang-gin-poc/entities"
	"github.com/pace-noge/golang-gin-poc/services"
)

type VideoController interface {
	FindAll() []entities.Video
	Save(ctx *gin.Context) entities.Video
}

type controller struct {
	service services.VideoService
}

func New(service services.VideoService) VideoController {
	return controller{
		service: service,
	}
}

func (c controller) FindAll() []entities.Video {
	return c.service.FindAll()
}

func (c controller) Save(ctx *gin.Context) entities.Video {
	var video entities.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
