package services

import (
	"fmt"
	"strconv"
	"unisun/api/unisun-authen-management-information/src/dao"
	"unisun/api/unisun-authen-management-information/src/models"
	"unisun/api/unisun-authen-management-information/src/ports/repositories"
)

type services struct {
	Repo repositories.RepositoriesUserAuthPermission
}

func NewUserAuthPermission(repo repositories.RepositoriesUserAuthPermission) *services {
	return &services{
		Repo: repo,
	}
}

func (srv *services) FindWithUser(userId string) (*models.UserAuthPermission, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, fmt.Errorf("%s{%v}", "Error step convert user ID from string to int.Error msg : ", err)
	}
	result := srv.Repo.FindbyUserid(id)
	body := dao.ConvertEntityToModel(&result)
	return body, nil
}

func (srv *services) CreateUserAuthPermission(payload models.UserAuthPermission) (bool, error) {
	body := dao.ConvertModelToEntity(&payload)
	srv.Repo.Create(*body)
	return true, nil
}

func (srv *services) UpdateUserAuthPermission(payload models.UserAuthPermission, tokenVersion string) (bool, error) {
	tkVersion, err := strconv.Atoi(tokenVersion)
	if err != nil {
		return false, fmt.Errorf("%s{%v}", "Error step convert token verion from string to int.Error msg : ", err)
	}
	body := dao.ConvertModelToEntity(&payload)
	srv.Repo.UpdateVersionToken(tkVersion, *body)
	return true, nil
}
