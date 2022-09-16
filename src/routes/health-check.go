package routes

import (
	"unisun/api/unisun-authen-management-information/src/ports/controllers"

	"github.com/gin-gonic/gin"
)

type controllerHealthCheck struct {
	Controller controllers.ControllerHealthCheck
}

func NewRouteHealthcheck(_controllerHealthCheck controllers.ControllerHealthCheck) *controllerHealthCheck {
	return &controllerHealthCheck{
		Controller: _controllerHealthCheck,
	}
}

func (srv *controllerHealthCheck) HealthCheck(g *gin.RouterGroup) {
	g.GET("/health-check", srv.Controller.HealthCheckHandler)
}
