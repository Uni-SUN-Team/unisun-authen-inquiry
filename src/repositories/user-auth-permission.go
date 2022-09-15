package repositories

import (
	"unisun/api/unisun-authen-inquiry/src/configs/database"
	"unisun/api/unisun-authen-inquiry/src/entities"

	"gorm.io/gorm"
)

type dbServices struct {
	Context *gorm.DB
}

func NewUserAuthPermission() *dbServices {
	return &dbServices{
		Context: database.DB,
	}
}

func (db *dbServices) FindbyUserid(userId int) entities.UserAuthPermission {
	user_permission := entities.UserAuthPermission{}
	db.Context.First(&user_permission, userId)
	return user_permission
}

func (db *dbServices) Create(Data entities.UserAuthPermission) {
	db.Context.Create(&Data)
}

func (db *dbServices) UpdateVersionToken(versionToken int, Data entities.UserAuthPermission) {
	db.Context.Model(&Data).Where("user_id", Data.UserId).Update("token_version", versionToken).Update("iat", Data.Iat).Update("ext", Data.Ext)
}
