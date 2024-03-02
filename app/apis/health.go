package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthAPIs(group *gin.RouterGroup, c *APIBootstrapContext) {
	group.Any("/", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"Status":  http.StatusOK,
				"Service": c.AppParam.ServiceName,
			},
		)
	})
}
