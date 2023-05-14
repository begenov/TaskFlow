package usercontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/begenov/TaskFlow/pkg/models"
	"github.com/gin-gonic/gin"
)

type res struct {
	Task []models.Todo `json:"successful"`
}

type aa struct {
	Res res `json:"successful"`
}

func (u *UserController) UserAllTasks(ctx *gin.Context) {
	resp, err := http.Get("http://localhost:8000/tasks")
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var task []models.Todo

	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		log.Fatal(string(body), err)
		return
	}

	ctx.JSON(http.StatusOK, task)
}
