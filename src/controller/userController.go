package controller

import (
	"time"

	"github.com/Bean-jun/Dlog/dao/entity"
	"github.com/Bean-jun/Dlog/pkg"
	"github.com/Bean-jun/Dlog/services"
	"github.com/Bean-jun/Dlog/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/*
	Register

TODO:
1. 返回用户信息
2. 接收更多参数
*/
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	userService := services.ImplUser(&services.UserService{})
	u, msg := userService.AddUser(username, password)
	if u.IsEmpty() {
		c.JSON(200, gin.H{
			"status": false,
			"msg":    msg,
			"data":   nil,
		})
	} else {
		token, _ := utils.GenerateToken(jwt.MapClaims{
			"exp": time.Now().Add(time.Second * time.Duration(pkg.Conf.Server.TokenExpire)).Unix(),
			"id":  u.ID,
		})

		responseUser := services.ResponseUser{}
		c.JSON(200, gin.H{
			"status": true,
			"msg":    "success",
			"data": map[string]interface{}{
				"token": token,
				"user":  responseUser.ToResponseUser(u),
			},
		})
	}
}

/*
	Login

TODO:
1. 账号密码加密处理
2. 账号添加锁号功能
3. 账号添加密码修改时间提醒
4. 添加验证码功能
*/
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	userService := services.ImplUser(&services.UserService{})
	u := userService.FindByUserName(username)

	if utils.CheckPasswordHash(u.Password, password) {
		token, _ := utils.GenerateToken(jwt.MapClaims{
			"exp": time.Now().Add(time.Second * time.Duration(pkg.Conf.Server.TokenExpire)).Unix(),
			"id":  u.ID,
		})
		responseUser := services.ResponseUser{}
		c.JSON(200, gin.H{
			"status": true,
			"msg":    "success",
			"data": map[string]interface{}{
				"token": token,
				"user":  responseUser.ToResponseUser(u),
			},
		})
	} else {
		c.JSON(200, gin.H{
			"status": false,
			"msg":    "账号或密码异常",
			"data":   nil,
		})
	}
}

/*
	GetUserInfo

TODO:
1. 用户敏感数据过滤
2. 用户额外数据处理
*/
func GetUserInfo(c *gin.Context) {
	if user, ok := c.Get("user"); !ok {
		c.Abort()
	} else {
		responseUser := services.ResponseUser{}
		c.JSON(200, responseUser.ToResponseUser(user.(entity.UserEntity)))
	}
}
