package usercontroller

import (
	"encoding/json"
	"io"

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
