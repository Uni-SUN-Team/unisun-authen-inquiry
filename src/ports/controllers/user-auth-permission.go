package controllers

import "github.com/gin-gonic/gin"

type ControllerUserAuthPermission interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}
