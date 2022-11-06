package middleware

import (
	"github.com/Bean-jun/Dlog/dao/entity"
	"github.com/Bean-jun/Dlog/services"
	"github.com/Bean-jun/Dlog/utils"
	"github.com/gin-gonic/gin"
)

func Auth() func(*gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		// 访客
		if token == "" {
			u := &entity.UserEntity{}
			c.Set("user", u)
		} else {
			// 网站用户
			parseToken, err := utils.ParseToken(token)
			// token 过期||异常
			if err != nil {
				utils.FalseResponse(c, "鉴权失败")
				c.Abort()
				return
			}

			uid := parseToken["id"]
			userService := services.ImplUser(&services.UserService{})
			u := userService.FindByUserID(int(uid.(float64)))
			// 用户已不存在
			if u == nil {
				utils.FalseResponse(c, "鉴权失败")
				c.Abort()
				return
			}
			c.Set("user", u)
		}

		c.Next()
	}
}
