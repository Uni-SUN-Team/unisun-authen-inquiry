package routes

import "github.com/gin-gonic/gin"

type RouteHealthCheck interface {
	HealthCheck(g *gin.RouterGroup)
}
