package controller

import (
	"github.com/begenov/TaskFlow/task-app/internal/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	services service.Tasks
}

func NewController(task service.Tasks) *Controller {
	return &Controller{
		services: task,
	}
}

func (c *Controller) Init(ctx *gin.Context) {
	router := gin.Default()

	task := router.Group("/tasks")

	{
		task.POST("/create", c.createTask)
		task.GET("", c.allTasks)
		task.GET("/:id", c.taskByID)
		task.PUT("/update", c.updateTask)
		task.DELETE("/delete", c.deleteTask)
	}
}
