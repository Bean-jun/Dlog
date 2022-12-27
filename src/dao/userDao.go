package dao

import (
	"time"

	"github.com/Bean-jun/Dlog/dao/entity"
	"github.com/Bean-jun/Dlog/utils"
)

func FindByUserName(username string) (*entity.UserEntity, error) {
	user := &entity.UserEntity{}
	if err := DB.First(user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func FindByUserID(id int) (*entity.UserEntity, error) {
	user := &entity.UserEntity{}
	if err := DB.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func AddUser(username, password string) (*entity.UserEntity, error) {
	_, hashPassword := utils.GeneratePasswordHash(password)
	user := &entity.UserEntity{
		Username:       username,
		Password:       hashPassword,
		ChangePassword: time.Now(),
		IsAdmin:        false,
		LoginAt:        time.Now(),
	}
	if err := DB.Select("Username", "Password", "ChangePassword", "IsAdmin", "LoginAt").
		Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
