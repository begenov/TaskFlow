package usercontroller

import (
	"fmt"
	"io"
	"net/http"

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
	fmt.Println(string(body))
}
