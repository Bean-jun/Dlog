package controller

import (
	"github.com/Bean-jun/Dlog/dao"
	"github.com/Bean-jun/Dlog/utils"
	"github.com/gin-gonic/gin"
)

// GetCaptcha @Schemes
// @Description 获取验证码
// @Tags GetCaptcha
// @Accept json
// @Produce json
// @Success 200 {object}  utils.Response
// @Router /api/v1/getCaptcha [get]
func GetCaptcha(c *gin.Context) {
	uuidStr, data := utils.GenerateCaptcha(dao.RDB)
	utils.TrueResponse(c, "success", gin.H{
		"code": uuidStr,
		"img":  data,
	})
}
