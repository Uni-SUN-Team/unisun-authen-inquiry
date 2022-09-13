package routes

import (
	"unisun/api/auth-processor-api/src/controllers"
	repositoies "unisun/api/auth-processor-api/src/repositories"
	"unisun/api/auth-processor-api/src/services"

	"github.com/gin-gonic/gin"
)

func UserAuthPermission(g *gin.RouterGroup) {
	repo := repositoies.NewUserAuthPermission()
	service := services.NewUserAuthPermission(repo)
	controller := controllers.NewUserAuthPermission(service)
	g.GET("/user-auth-permission/:id", controller.Get)
	g.POST("/user-auth-permission", controller.Create)
	g.PUT("/user-auth-permission/:version", controller.Update)
}
