package main

import (
	"log"

	"github.com/Bean-jun/Dlog/dao"

	"github.com/Bean-jun/Dlog/pkg"
	"github.com/Bean-jun/Dlog/router"
)

func InitEnv() {
	pkg.InitConfig("conf.yaml")
	dao.InitDB()
}

// @title Dlog
// @version 1.0
// @description Dlog后端api服务
// @contact.name Bean-jun
// @contact.url https://github.com/Bean-jun
// @contact.email 1342104001@qq.com
// @license.name MIT license
// @host 192.168.2.100:9090
// @BasePath /
func main() {
	InitEnv()
	route := router.InitRouter()
	log.Fatal(route.Run(pkg.Conf.Server.Port))
}
