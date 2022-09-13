package repositories

import "unisun/api/auth-processor-api/src/entities"

type RepositoriesUserAuthPermission interface {
	FindbyUserid(userId int) entities.UserAuthPermission
	Create(Data entities.UserAuthPermission)
	UpdateVersionToken(versionToken int, Data entities.UserAuthPermission)
}
