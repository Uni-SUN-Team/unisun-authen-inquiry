package routes

import (
	"unisun/api/unisun-authen-inquiry/src/controllers"

	"github.com/gin-gonic/gin"
)

func HealthCheck(g *gin.RouterGroup) {
	healController := controllers.NewControllerHealthCheckHandler()
	g.GET("/healcheck", healController.HealthCheckHandler)
}
