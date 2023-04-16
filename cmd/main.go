package main

import (
	"github.com/begenov/TaskFlow/cmd/controller"
	"github.com/begenov/TaskFlow/cmd/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VedioService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	mux := gin.Default()
	mux.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())

	})

	mux.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))

	})

	mux.Run(":8080")
}
