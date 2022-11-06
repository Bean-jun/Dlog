package router

import (
	_ "github.com/Bean-jun/Dlog/docs"

	"github.com/Bean-jun/Dlog/controller"
	"github.com/Bean-jun/Dlog/middleware"
	"github.com/Bean-jun/Dlog/pkg"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func getRouter() *gin.Engine {
	var engine *gin.Engine
	if pkg.Conf.Debug {
		engine = gin.Default()
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	} else {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
	}
	return engine
}

func InitRouter() *gin.Engine {
	e := getRouter()
	// Add route content
	api := e.Group("/api/v1")
	{
		api.POST("/login", controller.Login)
		api.POST("/register", controller.Register)
		api.GET("/userinfo", middleware.Auth(), controller.GetUserInfo)
		api.GET("/getCaptcha", controller.GetCaptcha)
	}
	return e
}
