package controller

import (
	"github.com/gin-gonic/gin"
)

func (c *controller) Router() *gin.Engine {

	mux := gin.Default()

	user := mux.Group("/user")
	{
		user.POST("/sign-up", c.user.SignUp)
		user.POST("/sign-in", c.user.SignIn)
		user.POST("/auth/refresh", c.user.UserRefresh)

		task := user.Group("/tasks")
		task.GET("", c.user.UserAllTasks)
	}

	return mux
}

/*
func (c *controller) homepage(ctx *gin.Context) {
	var data models.Data

	user_id, ok := ctx.Value("user_id").(int)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user ID"})
		return
	}
	user, err := c.service.User.UserByID(context.Background(), user_id)

	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"ERROR": "ERROR",
			})
		}
		return
	} else {
		data.UserName = user.Username
	}

	ctx.JSON(http.StatusOK, gin.H{
		"info": data,
	})

}
*/
