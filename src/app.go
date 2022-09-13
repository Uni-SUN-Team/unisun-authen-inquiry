package src

import (
	"strings"
	"unisun/api/auth-processor-api/docs"
	"unisun/api/auth-processor-api/src/configs/environment"
	"unisun/api/auth-processor-api/src/routes"

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
		routes.HealthCheck(g)
		routes.UserAuthPermission(g)
	}
	return r
}
