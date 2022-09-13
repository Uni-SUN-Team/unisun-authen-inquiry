package dao

import (
	"unisun/api/auth-processor-api/src/entities"
	"unisun/api/auth-processor-api/src/models"
)

func ConvertEntityToModel(payload *entities.UserAuthPermission) *models.UserAuthPermission {
	newPayload := models.UserAuthPermission{}
	newPayload.Ext = payload.Ext
	newPayload.Iat = payload.Iat
	newPayload.TokenVersion = payload.TokenVersion
	newPayload.UserId = payload.UserId
	return &newPayload
}

func ConvertModelToEntity(payload *models.UserAuthPermission) *entities.UserAuthPermission {
	newPayload := entities.UserAuthPermission{}
	newPayload.Ext = payload.Ext
	newPayload.Iat = payload.Iat
	newPayload.TokenVersion = payload.TokenVersion
	newPayload.UserId = payload.UserId
	return &newPayload
}
