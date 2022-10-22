package services

import (
	"github.com/Bean-jun/Dlog/dao"
	"github.com/Bean-jun/Dlog/dao/entity"
)

type UserService struct {
}

func (u *UserService) FindByUserName(username string) entity.UserEntity {
	user, err := dao.FindByUserName(username)
	if err != nil {
		return entity.UserEntity{}
	}
	return user
}

func (u *UserService) FindByUserID(id int) entity.UserEntity {
	user, err := dao.FindByUserID(id)
	if err != nil {
		return entity.UserEntity{}
	}
	return user
}

func (u *UserService) AddUser(username, password string) (entity.UserEntity, string) {
	user := u.FindByUserName(username)
	if !user.IsEmpty() {
		return entity.UserEntity{}, "账号已注册"
	}
	addUser, err := dao.AddUser(username, password)
	if err != nil {
		return entity.UserEntity{}, "账号注册失败"
	}
	return addUser, "注册成功"
}
