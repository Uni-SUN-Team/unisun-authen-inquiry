package routes

import "github.com/gin-gonic/gin"

type RouteUserAuthPermission interface {
	UserAuthPermission(g *gin.RouterGroup)
}
