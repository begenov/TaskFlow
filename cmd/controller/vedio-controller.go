package controller

import (
	"fmt"

	"github.com/begenov/TaskFlow/cmd/entity"
	"github.com/begenov/TaskFlow/cmd/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}
type controller struct {
	service service.VedioService
}

func New(service service.VedioService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()

}
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	if err := ctx.Bind(&video); err != nil {
		fmt.Print(err)
	}
	c.service.Save(video)
	return video
}
