package services

import (
	"time"

	"github.com/Bean-jun/Dlog/dao/entity"
)

type ImplUser interface {
	FindByUserName(string) *entity.UserEntity
	FindByUserID(int) *entity.UserEntity
	AddUser(string, string) (*entity.UserEntity, string)
}

// ResponseUser 响应结构体
type ResponseUser struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
}

func (u *ResponseUser) ToResponseUser(userEntity *entity.UserEntity) ResponseUser {
	return ResponseUser{
		Id:        userEntity.ID,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		Gender:    userEntity.Gender,
		Phone:     userEntity.Phone,
	}
}
