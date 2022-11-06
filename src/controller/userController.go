package controller

import (
	"fmt"
	"time"

	"github.com/Bean-jun/Dlog/dao"
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
	username := utils.RsaDecryptFactory(c.PostForm("username"), pkg.Conf.Server.Cert.PrivateKey)
	password := utils.RsaDecryptFactory(c.PostForm("password"), pkg.Conf.Server.Cert.PrivateKey)

	userService := services.ImplUser(&services.UserService{})
	u, msg := userService.AddUser(username, password)
	if u == nil {
		utils.FalseResponse(c, msg)
		return
	}

	token, _ := utils.GenerateToken(jwt.MapClaims{
		"exp": time.Now().Add(time.Second * time.Duration(pkg.Conf.Server.TokenExpire)).Unix(),
		"id":  u.ID,
	})
	responseUser := services.ResponseUser{}
	utils.TrueResponse(c, "success", map[string]interface{}{
		"token": token,
		"user":  responseUser.ToResponseUser(u),
	})
}

/*
	Login

TODO:
1. 账号密码加密处理 √
2. 账号添加锁号功能
3. 账号添加密码修改时间提醒
4. 添加验证码功能
*/

// Login @Schemes
// @Description 用户登录
// @Tags Login
// @Accept mpfd
// @Produce json
// @Param   username     formData    string     true  "Username"
// @Param   password     formData    string     true  "Password"
// @Success 200 {object}  utils.Response
// @Router /api/v1/login [post]
func Login(c *gin.Context) {
	username := utils.RsaDecryptFactory(c.PostForm("username"), pkg.Conf.Server.Cert.PrivateKey)
	password := utils.RsaDecryptFactory(c.PostForm("password"), pkg.Conf.Server.Cert.PrivateKey)

	userService := services.ImplUser(&services.UserService{})
	u := userService.FindByUserName(username)
	if u == nil {
		utils.FalseResponse(c, "账户不存在")
		return
	}

	msg := "success"

	// 校验账号是否长时间未修改密码
	if pkg.Conf.Account.AcountMaxModifyPasswordInterval != -1 {
		if u.ChangePassword.IsZero() {
			// 更新用户修改密码时间
			u.ChangePassword = time.Now()
			dao.DB.Save(u)
		}
		if u.ChangePassword.Add(time.Second * time.Duration(pkg.Conf.Account.AcountMaxNotActive)).Before(time.Now()) {
			interval := int64(time.Since(u.ChangePassword).Hours() / 24)
			msg = fmt.Sprintf("success$您已经%d天未修改密码,为保障您的账户安全,请您及时修改密码", interval)
		}
	}

	// 校验是否符合登录条件
	if pkg.Conf.Account.AcountMaxNotActive != -1 {
		if u.LoginAt.IsZero() {
			// 更新登录时间
			u.LoginAt = time.Now()
			dao.DB.Save(u)
		}
		if u.LoginAt.Add(time.Second * time.Duration(pkg.Conf.Account.AcountMaxNotActive)).Before(time.Now()) {
			utils.FalseResponse(c, "由于您长时间未登录您的账户,您的账号已被冻结,请联系管理员处理!")
			return
		}
	}

	// 校验账号是否存在被锁情况
	if u.ErrNum >= uint(pkg.Conf.Account.LoginErrNum) {
		if u.ErrAt.IsZero() {
			// 设置错误最新时间
			u.ErrAt = time.Now()
			dao.DB.Save(u)
		}
		if time.Now().Add(-1 * time.Second * time.Duration(pkg.Conf.Account.LoginErrLock)).Before(u.ErrAt) {
			interval := int64(u.ErrAt.Sub(time.Now().Add(-1 * time.Second * time.Duration(pkg.Conf.Account.LoginErrLock))).Minutes())
			msg = fmt.Sprintf("由于您多次输入错误的账号或密码,账号已被锁定,请%d分钟后再次尝试!", interval)
			utils.FalseResponse(c, msg)
			return
		}
	}

	// 校验密码 密码不成功,锁账号机制触发
	if !utils.CheckPasswordHash(u.Password, password) {
		if u.ErrAt.IsZero() {
			u.ErrAt = time.Now()
		}
		if time.Now().Add(-1 * time.Second * time.Duration(pkg.Conf.Account.LoginErrInterval)).Before(u.ErrAt) {
			u.ErrNum++
		} else {
			u.ErrNum = 0
		}

		// 更新用户密码尝试次数&错误时间
		dao.DB.Save(u)

		msg = "账号或密码异常"
		if u.ErrNum >= uint(pkg.Conf.Account.LoginErrTips) {
			nums := uint(pkg.Conf.Account.LoginErrNum) - u.ErrNum
			if nums > 0 {
				msg = fmt.Sprintf("您已输入%d次错误的账号或密码,您还有%d次尝试机会!", u.ErrNum, nums)
			} else {
				msg = fmt.Sprintf("您已输入%d次错误的账号或密码,账号已被锁定,请%d分钟之后再次尝试!", pkg.Conf.Account.LoginErrNum, pkg.Conf.Account.LoginErrLock/60)
			}
		}
		utils.FalseResponse(c, msg)
		return
	}

	token, _ := utils.GenerateToken(jwt.MapClaims{
		"exp": time.Now().Add(time.Second * time.Duration(pkg.Conf.Server.TokenExpire)).Unix(),
		"id":  u.ID,
	})

	// 更新用户登录时间
	u.LoginAt = time.Now()
	dao.DB.Save(u)

	responseUser := services.ResponseUser{}
	utils.TrueResponse(c, msg, map[string]interface{}{
		"token": token,
		"user":  responseUser.ToResponseUser(u),
	})
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
		utils.TrueResponse(c, "success", responseUser.ToResponseUser(user.(*entity.UserEntity)))
	}
}
