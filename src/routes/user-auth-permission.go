package routes

import (
	"unisun/api/unisun-authen-inquiry/src/controllers"
	"unisun/api/unisun-authen-inquiry/src/repositories"
	"unisun/api/unisun-authen-inquiry/src/services"

	"github.com/gin-gonic/gin"
)

func UserAuthPermission(g *gin.RouterGroup) {
	repo := repositories.NewUserAuthPermission()
	service := services.NewUserAuthPermission(repo)
	controller := controllers.NewUserAuthPermission(service)
	g.GET("/user-auth-permission/:id", controller.Get)
	g.POST("/user-auth-permission", controller.Create)
	g.PUT("/user-auth-permission/:version", controller.Update)
}
