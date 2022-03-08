package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pace-noge/golang-gin-poc/entities"
	"github.com/pace-noge/golang-gin-poc/services"
	"github.com/pace-noge/golang-gin-poc/validators"
)

type VideoController interface {
	FindAll() []entities.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service services.VideoService
}

var validate *validator.Validate

func New(service services.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return controller{
		service: service,
	}
}

func (c controller) FindAll() []entities.Video {
	return c.service.FindAll()
}

func (c controller) Save(ctx *gin.Context) error {
	var video entities.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":   "Video Page",
		"results": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
