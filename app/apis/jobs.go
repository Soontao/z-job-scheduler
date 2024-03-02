package apis

import (
	"net/http"

	"fornever.org/app/model"
	"github.com/gin-gonic/gin"
)

func JobAPIs(group *gin.RouterGroup, c *APIBootstrapContext) {
	group.GET("/tasks", func(ctx *gin.Context) {
		var tasks []model.Task
		// TODO: limit, order by ....
		c.DB.Find(&tasks)
		ctx.JSON(http.StatusOK, gin.H{"results": tasks})
	})
	group.POST("/tasks", func(ctx *gin.Context) {
		var payload model.Task
		ctx.BindJSON(&payload)
		c.DB.Save(&payload)
		ctx.JSON(http.StatusAccepted, gin.H{"results": payload})
	})
}
