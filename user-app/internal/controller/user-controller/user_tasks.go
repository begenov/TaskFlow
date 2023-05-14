package usercontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"log"
	"net/http"

	"github.com/begenov/TaskFlow/pkg/models"
	"github.com/gin-gonic/gin"
)

func (u *UserController) UserAllTasks(ctx *gin.Context) {
	resp, err := http.Get("http://localhost:8000/tasks")
	if err != nil {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var task []models.Todo

	if err := json.Unmarshal(body, &task); err != nil {
		log.Fatal(string(body), err)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (u *UserController) UserCreateTask(ctx *gin.Context) {
	user_id, ok := ctx.Value("user_id").(int)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user ID"})
		return
	}

	var inputTask models.Todo

	if err := ctx.BindJSON(&inputTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": fmt.Errorf("%w", err),
		})
		return
	}

	body, err := json.Marshal(inputTask)

	if err != nil {
		return
	}

	resp, err := http.Post("http://localhost:8000/tasks/create/"+strconv.Itoa(user_id), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, resp.Body)

}