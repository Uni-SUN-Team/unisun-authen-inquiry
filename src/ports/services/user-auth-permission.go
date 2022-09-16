package services

import "unisun/api/unisun-authen-management-information/src/models"

type ServiceUserAuthPermission interface {
	FindWithUser(userId string) (*models.UserAuthPermission, error)
	CreateUserAuthPermission(payload models.UserAuthPermission) (bool, error)
	UpdateUserAuthPermission(payload models.UserAuthPermission, tokenVersion string) (bool, error)
}
