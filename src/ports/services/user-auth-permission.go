package services

import "unisun/api/auth-processor-api/src/models"

type ServiceUserAuthPermission interface {
	FindWithUser(userId string) (*models.UserAuthPermission, error)
	CreateUserAuthPermission(payload models.UserAuthPermission) (bool, error)
	UpdateUserAuthPermission(payload models.UserAuthPermission, tokenVersion string) (bool, error)
}
