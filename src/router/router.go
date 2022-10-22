package router

import (
	"github.com/Bean-jun/Dlog/pkg"
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	var engine *gin.Engine
	if pkg.Conf.Debug {
		engine = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
	}
	return engine
}

func InitRouter() *gin.Engine {
	e := getRouter()
	// Add route content
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": true,
			"msg":    "success",
			"code":   200,
		})
	})
	return e
}
