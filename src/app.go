package src

import (
	"strings"
	"unisun/api/unisun-authen-management-information/docs"
	"unisun/api/unisun-authen-management-information/src/configs/environment"
	"unisun/api/unisun-authen-management-information/src/controllers"
	"unisun/api/unisun-authen-management-information/src/repositories"
	"unisun/api/unisun-authen-management-information/src/routes"
	"unisun/api/unisun-authen-management-information/src/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @termsOfService 		http://swagger.io/terms/
// @contact.name 		API Support
// @contact.url 		http://www.swagger.io/support
// @contact.email 		support@swagger.io

// @license.name 		MIT License Copyright (c) 2022 Uni-SUN-Team
// @license.url 		${host}/${project_name}/api/v1/license

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func App() *gin.Engine {
	appEnv := environment.ENV.App
	ginEnv := environment.ENV.Gin
	swagEnv := environment.ENV.Swag
	docs.SwaggerInfo.Title = swagEnv.Title
	docs.SwaggerInfo.Description = swagEnv.Description
	docs.SwaggerInfo.Version = swagEnv.Version
	docs.SwaggerInfo.Host = swagEnv.Host
	docs.SwaggerInfo.BasePath = strings.Join([]string{appEnv.ContextPath, ginEnv.RootPath, ginEnv.Version}, "/")
	docs.SwaggerInfo.Schemes = strings.Split(swagEnv.Schemes, ",")

	r := gin.Default()
	r.SetTrustedProxies(strings.Split(ginEnv.Configs.TrustedProxies, ","))
	g := r.Group(strings.Join([]string{appEnv.ContextPath, ginEnv.RootPath, ginEnv.Version}, "/"))
	{
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		g.StaticFile("/license", "./LICENSE")
		mapRouteHealthCheck(g)
		mapRouteUserAuthPermission(g)
	}
	return r
}

func mapRouteHealthCheck(g *gin.RouterGroup) {
	controller := controllers.NewControllerHealthCheckHandler()
	route := routes.NewRouteHealthcheck(controller)
	route.HealthCheck(g)
}

func mapRouteUserAuthPermission(g *gin.RouterGroup) {
	repository := repositories.NewUserAuthPermission()
	service := services.NewUserAuthPermission(repository)
	controller := controllers.NewUserAuthPermission(service)
	route := routes.NewRouteUserAuthPermission(controller)
	route.UserAuthPermission(g)
}
