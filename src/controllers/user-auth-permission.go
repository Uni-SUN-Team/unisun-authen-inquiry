package controllers

import (
	"fmt"
	"unisun/api/unisun-authen-inquiry/src/models"
	"unisun/api/unisun-authen-inquiry/src/ports/services"

	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/handlercontrol/response/constants"
	serviceResponse "github.com/narawichsaphimarn/handlercontrol/response/services"
)

type service struct {
	Service services.ServiceUserAuthPermission
}

func NewUserAuthPermission(ser services.ServiceUserAuthPermission) *service {
	return &service{
		Service: ser,
	}
}

func (srv *service) Get(c *gin.Context) {
	id := c.Param("id")
	result, err := srv.Service.FindWithUser(id)
	if err != nil {
		code := constants.InternalServerError
		msg := serviceResponse.HttpMessage(code)
		c.AbortWithStatusJSON(code, serviceResponse.MapResponseUnsuccess(code, msg, err.Error(), nil))
	}
	code := constants.OK
	msg := serviceResponse.HttpMessage(code)
	c.AbortWithStatusJSON(code, serviceResponse.MapResponseSuccess(code, msg, result))
}

func (srv *service) Create(c *gin.Context) {
	body := models.UserAuthPermission{}
	if err := c.ShouldBindJSON(&body); err != nil {
		code := constants.InternalServerError
		msg := serviceResponse.HttpMessage(code)
		c.AbortWithStatusJSON(code, serviceResponse.MapResponseUnsuccess(code, msg, err.Error(), nil))
	}
	status, err := srv.Service.CreateUserAuthPermission(body)
	if err != nil {
		code := constants.InternalServerError
		msg := serviceResponse.HttpMessage(code)
		c.AbortWithStatusJSON(code, serviceResponse.MapResponseUnsuccess(code, msg, err.Error(), body))
	}
	if status {
		code := constants.OK
		msg := serviceResponse.HttpMessage(code)
		c.AbortWithStatusJSON(code, serviceResponse.MapResponseSuccess(code, msg, nil))
	} else {
		code := constants.InternalServerError
		msg := serviceResponse.HttpMessage(code)
		c.AbortWithStatusJSON(code, serviceResponse.MapResponseUnsuccess(code, msg, fmt.Errorf("%s%v", "Error status is ", status).Error(), body))
	}
}

func (srv *service) Update(c *gin.Context) {
	version := c.Param("version")
	body := models.UserAuthPermission{}
	if err := c.ShouldBindJSON(&body); err != nil {
		code := constants.InternalServerError
		msg := serviceResponse.HttpMessage(code)
		c.AbortWithStatusJSON(code, serviceResponse.MapResponseUnsuccess(code, msg, err.Error(), nil))
	}
	status, err := srv.Service.UpdateUserAuthPermission(body, version)
	if err != nil {
		code := constants.InternalServerError
		msg := serviceResponse.HttpMessage(code)
		c.AbortWithStatusJSON(code, serviceResponse.MapResponseUnsuccess(code, msg, err.Error(), body))
	}
	if status {
		code := constants.OK
		msg := serviceResponse.HttpMessage(code)
		c.AbortWithStatusJSON(code, serviceResponse.MapResponseSuccess(code, msg, nil))
	} else {
		code := constants.InternalServerError
		msg := serviceResponse.HttpMessage(code)
		c.AbortWithStatusJSON(code, serviceResponse.MapResponseUnsuccess(code, msg, fmt.Errorf("%s%v", "Error status is ", status).Error(), body))
	}
}
