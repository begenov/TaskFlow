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

func (c *Controller) Init() *gin.Engine {
	router := gin.Default()

	task := router.Group("/tasks")

	{
		task.GET("", c.allTasks)
		task.GET("/:id", c.taskByID)
		task.POST("/create/:userID", c.createTask)
		task.PUT("/update", c.updateTask)
		task.DELETE("/:id", c.deleteTask)
	}

	return router
}
