package routes

import (
	"unisun/api/unisun-authen-management-information/src/ports/controllers"

	"github.com/gin-gonic/gin"
)

type controllerUserAuthPermission struct {
	Controller controllers.ControllerUserAuthPermission
}

func NewRouteUserAuthPermission(_controllerUserAuthPermission controllers.ControllerUserAuthPermission) *controllerUserAuthPermission {
	return &controllerUserAuthPermission{
		Controller: _controllerUserAuthPermission,
	}
}

func (srv *controllerUserAuthPermission) UserAuthPermission(g *gin.RouterGroup) {
	g.GET("/user-auth-permission/:id", srv.Controller.Get)
	g.POST("/user-auth-permission", srv.Controller.Create)
	g.PUT("/user-auth-permission/:version", srv.Controller.Update)
}
