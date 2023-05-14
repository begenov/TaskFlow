package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/begenov/TaskFlow/pkg/models"
	"github.com/gin-gonic/gin"
)

func (c *Controller) createTask(ctx *gin.Context) {
	var inputTask models.Todo

	userID, ok := ctx.Value("user_id").(int)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user ID"})
		return
	}

	// userID := 1
	if err := ctx.BindJSON(&inputTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": fmt.Errorf("%w", err),
		})
		return
	}

	inputTask.CreatedAt = time.Now()
	inputTask.ID = userID

	if err := c.services.CreateTask(context.Background(), inputTask); err != nil {
		log.Println("errr", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": fmt.Errorf("%w", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"successful": "successful: create task",
	})

}

func (c *Controller) allTasks(ctx *gin.Context) {
	tasks, err := c.services.AllTasks(context.Background())
	if err != nil {
		fmt.Println(err, "errrrrrrrrrrr")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": fmt.Errorf("%w", err),
		})
		return
	}

	ctx.JSON(http.StatusOK,
		tasks,
	)
}

func (c *Controller) taskByID(ctx *gin.Context) {

	task_id := ctx.Param("id")

	id, err := strconv.Atoi(task_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": fmt.Errorf("%w", err),
		})
		return
	}

	task, err := c.services.TaskByID(context.Background(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": fmt.Errorf("%w", err),
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"successful": task,
	})

}

func (c *Controller) updateTask(ctx *gin.Context) {
	var inputTask models.Todo

	if err := ctx.BindJSON(&inputTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": fmt.Errorf("%w", err),
		})
		return
	}

	if err := c.services.UpdateTask(context.Background(), inputTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": fmt.Errorf("%w", err),
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"successful": "successful",
	})

}

func (c *Controller) deleteTask(ctx *gin.Context) {

	taskID := ctx.Param("id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": fmt.Errorf("%w", err),
		})
		return
	}

	if err := c.services.DeleteTask(context.Background(), id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": fmt.Errorf("%w", err),
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"successful": "successful",
	})
}
