package routes

import (
	"unisun/api/auth-processor-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func HealthCheck(g *gin.RouterGroup) {
	healController := controllers.NewControllerHealthCheckHandler()
	g.GET("/healcheck", healController.HealthCheckHandler)
}
