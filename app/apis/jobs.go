package apis

import "github.com/gin-gonic/gin"

func JobAPIs(group *gin.RouterGroup, c *APIBootstrapContext) {
	group.GET("/task", func(ctx *gin.Context) {

	})
	group.POST("/task", func(ctx *gin.Context) {

	})
}
