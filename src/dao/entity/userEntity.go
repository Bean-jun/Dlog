package entity

import (
	"reflect"
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	Username       string
	Email          string
	Gender         string
	Phone          string
	Password       string
	IsAdmin        bool
	ChangePassword time.Time
	LoginAt        time.Time // 登录时间用于处理后续异地登录问题
	ErrAt          time.Time // 密码输入错误时间
	ErrNum         uint      //错误密码次数
}

func (u *UserEntity) TableName() string {
	return "users"
}

func (u UserEntity) IsEmpty() bool {
	return reflect.DeepEqual(u, UserEntity{})
}
